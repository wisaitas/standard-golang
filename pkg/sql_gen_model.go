package pkg

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

// SQLToModelConverter converts SQL table definitions to Go model structs
type SQLToModelConverter struct {
	SourcePath      string
	DestinationPath string
}

// ForeignKeyInfo stores information about a foreign key relationship
type ForeignKeyInfo struct {
	ColumnName       string
	ReferencedTable  string
	ReferencedColumn string
}

// TableInfo stores information about a table
type TableInfo struct {
	Name        string
	ColumnNames []string
	ColumnTypes []string
	ColumnTags  []string
	ForeignKeys []ForeignKeyInfo
	EnumTypes   map[string]bool
}

// NewSQLToModelConverter creates a new SQL to model converter
func NewSQLToModelConverter(sourcePath, destinationPath string) *SQLToModelConverter {
	return &SQLToModelConverter{
		SourcePath:      sourcePath,
		DestinationPath: destinationPath,
	}
}

// sqlTypeToGoType maps SQL types to Go types
func sqlTypeToGoType(sqlType string, isRequired bool, isEnum ...bool) string {
	// Check if isEnum is provided and set the enumFlag accordingly
	enumFlag := false
	if len(isEnum) > 0 {
		enumFlag = isEnum[0]
	}

	sqlType = strings.ToLower(sqlType)

	if enumFlag || strings.Contains(sqlType, "enum") {
		if isRequired {
			return "string"
		}
		return "*string"
	} else if strings.Contains(sqlType, "uuid") {
		return "uuid.UUID"
	} else if strings.Contains(sqlType, "varchar") || strings.Contains(sqlType, "text") {
		if isRequired {
			return "string"
		}
		return "*string"
	} else if strings.Contains(sqlType, "integer") || strings.Contains(sqlType, "int") {
		if isRequired {
			return "int"
		}
		return "*int"
	} else if strings.Contains(sqlType, "decimal") || strings.Contains(sqlType, "numeric") ||
		strings.Contains(sqlType, "float") || strings.Contains(sqlType, "double") {
		if isRequired {
			return "float64"
		}
		return "*float64"
	} else if strings.Contains(sqlType, "bool") {
		if isRequired {
			return "bool"
		}
		return "*bool"
	} else if strings.Contains(sqlType, "timestamp") {
		if isRequired {
			return "time.Time"
		}
		return "*time.Time"
	} else if strings.Contains(sqlType, "date") {
		if isRequired {
			return "time.Time"
		}
		return "*time.Time"
	} else if strings.Contains(sqlType, "json") || strings.Contains(sqlType, "jsonb") {
		return "json.RawMessage"
	}

	return "interface{}"
}

// snakeToCamel converts snake_case to CamelCase
func snakeToCamel(s string) string {
	words := strings.Split(s, "_")
	for i := range words {
		if len(words[i]) > 0 {
			r := []rune(words[i])
			r[0] = unicode.ToUpper(r[0])
			words[i] = string(r)
		}
	}
	return strings.Join(words, "")
}

// singularize attempts to convert plural table names to singular for model names with proper handling
// of common edge cases
func singularize(s string) string {
	// Define singularization rules from most specific to most general
	rules := []struct {
		suffix      string
		replacement string
	}{
		// Special cases
		{"ies", "y"},              // e.g., "cities" -> "city"
		{"provinces", "province"}, // Handle province correctly
		{"statuses", "status"},    // Handle status correctly
		{"addresses", "address"},  // Handle address correctly
		{"analyses", "analysis"},  // Handle analysis correctly
		{"series", "series"},      // Handle series (unchanged)
		{"species", "species"},    // Handle species (unchanged)

		// General rules (applied if no specific rule matches)
		{"xes", "x"},   // e.g., "boxes" -> "box"
		{"ches", "ch"}, // e.g., "beaches" -> "beach"
		{"shes", "sh"}, // e.g., "dishes" -> "dish"
		{"zzes", "z"},  // e.g., "quizzes" -> "quiz"
		{"sses", "ss"}, // e.g., "classes" -> "class"
		{"ves", "f"},   // e.g., "wolves" -> "wolf"
		{"s", ""},      // Simple plural removal (most common case)
	}

	// Try to apply rules until one matches
	lowerS := strings.ToLower(s)
	for _, rule := range rules {
		if strings.HasSuffix(lowerS, rule.suffix) {
			// Preserve original case by getting the non-suffix part
			prefix := s[:len(s)-len(rule.suffix)]
			// For simple 's' removal, check that we're not dealing with a non-plural word ending in 's'
			if rule.suffix == "s" && rule.replacement == "" {
				// Check for common non-plural words ending in 's'
				nonPlurals := []string{"status", "campus", "bus", "virus"}
				for _, np := range nonPlurals {
					if lowerS == np {
						return s // Return unchanged
					}
				}
			}
			return prefix + rule.replacement
		}
	}

	// If no rule matched, return the original string
	return s
}

// pluralize makes a word plural - used for relationship field names
func pluralize(s string) string {
	// Special cases
	specialCases := map[string]string{
		"Address": "Addresses",
		"Status":  "Statuses",
		"Person":  "People",
		"Child":   "Children",
		"Foot":    "Feet",
		"Tooth":   "Teeth",
		"Goose":   "Geese",
		"Mouse":   "Mice",
		"Man":     "Men",
		"Woman":   "Women",
		"Ox":      "Oxen",
	}

	if plural, ok := specialCases[s]; ok {
		return plural
	}

	// General pluralization rules
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "x") || strings.HasSuffix(s, "z") ||
		strings.HasSuffix(s, "ch") || strings.HasSuffix(s, "sh") {
		return s + "es"
	} else if strings.HasSuffix(s, "y") && len(s) > 1 {
		// If the word ends in y preceded by a consonant, change y to ies
		c := s[len(s)-2]
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' {
			return s[:len(s)-1] + "ies"
		}
	}

	return s + "s"
}

// generateModelName generates a proper model name from a table name
func generateModelName(tableName string) string {
	// First singularize the table name
	singular := singularize(tableName)

	// Then convert to camel case
	modelName := snakeToCamel(singular)

	return modelName
}

// formatIDField properly formats an ID field (e.g., province_id becomes ProvinceID)
func formatIDField(fieldName string) string {
	if strings.HasSuffix(fieldName, "_id") {
		baseName := fieldName[:len(fieldName)-3]
		return snakeToCamel(baseName) + "ID"
	}
	return snakeToCamel(fieldName)
}

// getSQLFiles returns all SQL files in the source directory
func (c *SQLToModelConverter) getSQLFiles() ([]string, error) {
	var files []string

	err := filepath.Walk(c.SourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".sql") {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// extractTableInfo extracts table name, columns, and foreign keys from SQL CREATE TABLE statement
func extractTableInfo(sqlContent string) (*TableInfo, error) {
	// Create a map to store enum types
	enumTypes := make(map[string]bool)

	// Find all enum declarations in the SQL content
	enumRegex := regexp.MustCompile(`CREATE\s+TYPE\s+([a-zA-Z0-9_]+)\s+AS\s+ENUM`)
	enumMatches := enumRegex.FindAllStringSubmatch(sqlContent, -1)
	for _, match := range enumMatches {
		if len(match) >= 2 {
			enumTypes[match[1]] = true
		}
	}

	// Extract table name
	tableNameRegex := regexp.MustCompile(`CREATE\s+TABLE\s+IF\s+NOT\s+EXISTS\s+([a-zA-Z0-9_]+)\s*\(`)
	matches := tableNameRegex.FindStringSubmatch(sqlContent)
	if len(matches) < 2 {
		return nil, fmt.Errorf("table name not found")
	}
	tableName := matches[1]

	// Create table info
	tableInfo := &TableInfo{
		Name:      tableName,
		EnumTypes: enumTypes,
	}

	// Extract columns
	scanner := bufio.NewScanner(strings.NewReader(sqlContent))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "--") || line == "" || strings.HasPrefix(line, "CREATE TABLE") ||
			strings.HasPrefix(line, ");") || strings.HasPrefix(line, ")") || strings.HasPrefix(line, "CREATE TYPE") {
			continue
		}

		// Extract foreign key constraints
		if strings.HasPrefix(line, "ALTER TABLE") && strings.Contains(line, "FOREIGN KEY") {
			// Parse foreign key information
			fkRegex := regexp.MustCompile(`ALTER\s+TABLE\s+[a-zA-Z0-9_]+\s+ADD\s+CONSTRAINT\s+[a-zA-Z0-9_]+\s+FOREIGN\s+KEY\s+\(([a-zA-Z0-9_]+)\)\s+REFERENCES\s+([a-zA-Z0-9_]+)\(([a-zA-Z0-9_]+)\)`)
			fkMatches := fkRegex.FindStringSubmatch(line)

			if len(fkMatches) >= 4 {
				fk := ForeignKeyInfo{
					ColumnName:       fkMatches[1],
					ReferencedTable:  fkMatches[2],
					ReferencedColumn: fkMatches[3],
				}
				tableInfo.ForeignKeys = append(tableInfo.ForeignKeys, fk)
			}

			continue
		}

		// Remove trailing comma
		if strings.HasSuffix(line, ",") {
			line = line[:len(line)-1]
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		columnName := strings.Trim(parts[0], `"`)

		// Skip if this is a constraint or primary key definition
		if columnName == "PRIMARY" || columnName == "CONSTRAINT" || columnName == "FOREIGN" {
			continue
		}

		columnType := parts[1]

		// Check if this is an enum type
		isEnum := enumTypes[columnType] || strings.Contains(strings.ToLower(columnType), "enum")

		// Build GORM tag - without backticks, they'll be added when writing to file
		tag := fmt.Sprintf(`gorm:"type:%s`, columnType)

		// Check for NOT NULL
		isRequired := strings.Contains(strings.ToUpper(line), "NOT NULL")
		if isRequired {
			tag += ";not null"
		}

		// Check for UNIQUE
		if strings.Contains(strings.ToUpper(line), "UNIQUE") {
			tag += ";unique"
		}

		// Check for DEFAULT
		defaultRegex := regexp.MustCompile(`DEFAULT\s+([^,\s]+)`)
		defaultMatches := defaultRegex.FindStringSubmatch(line)
		if len(defaultMatches) > 1 {
			tag += fmt.Sprintf(`;default:%s`, defaultMatches[1])
		}

		// Add primary key if this is the ID column
		if columnName == "id" {
			tag += ";primaryKey"
		}

		// If it's an enum, add a comment to the GORM tag for clarity
		if isEnum {
			tag += ";comment:enum type"
		}

		tag += `"`

		// Skip base model fields
		if columnName == "id" || columnName == "version" ||
			columnName == "created_at" || columnName == "created_by" ||
			columnName == "updated_at" || columnName == "updated_by" ||
			columnName == "deleted_at" {
			continue
		}

		tableInfo.ColumnNames = append(tableInfo.ColumnNames, columnName)
		tableInfo.ColumnTypes = append(tableInfo.ColumnTypes, columnType)
		tableInfo.ColumnTags = append(tableInfo.ColumnTags, tag)
	}

	return tableInfo, nil
}

// getAllTableInfos extracts all table info from all SQL files
func (c *SQLToModelConverter) getAllTableInfos() (map[string]*TableInfo, error) {
	sqlFiles, err := c.getSQLFiles()
	if err != nil {
		return nil, err
	}

	tableInfos := make(map[string]*TableInfo)

	for _, sqlFile := range sqlFiles {
		// Skip seed files
		if strings.Contains(sqlFile, "seed-") {
			continue
		}

		// Read SQL file
		sqlContent, err := os.ReadFile(sqlFile)
		if err != nil {
			return nil, err
		}

		// Extract table info
		tableInfo, err := extractTableInfo(string(sqlContent))
		if err != nil {
			fmt.Printf("Error extracting table info from %s: %v\n", sqlFile, err)
			continue
		}

		tableInfos[tableInfo.Name] = tableInfo
	}

	return tableInfos, nil
}

// generateModelFile generates a Go model file from table info
func (c *SQLToModelConverter) generateModelFile(tableInfo *TableInfo, allTableInfos map[string]*TableInfo) error {
	// Generate proper model name and file name
	modelName := generateModelName(tableInfo.Name)

	// Get the singularized table name for the file name
	tableSingular := singularize(tableInfo.Name)

	// Create model file path - use the properly singularized name
	modelFileName := fmt.Sprintf("%s.go", tableSingular)
	modelFilePath := filepath.Join(c.DestinationPath, modelFileName)

	// Create model file
	modelFile, err := os.Create(modelFilePath)
	if err != nil {
		return err
	}
	defer modelFile.Close()

	// Write model file
	writer := bufio.NewWriter(modelFile)

	// Package declaration
	packageName := filepath.Base(c.DestinationPath)
	// Add package comment stating the file is auto-generated
	fmt.Fprintf(writer, "// Code generated by SQL-to-Model generator.\n")
	fmt.Fprintf(writer, "package %s\n\n", packageName)

	// Track imports needed
	needsTime := true // Always need time for BaseModel fields
	needsJSON := false
	needsUUID := true // Always need UUID for BaseModel fields
	needsGorm := true // Always need gorm for BaseModel fields and BeforeUpdate

	// Check for required imports
	for _, colType := range tableInfo.ColumnTypes {
		if strings.Contains(strings.ToLower(colType), "json") {
			needsJSON = true
		}
	}

	// Check if we have foreign keys
	if len(tableInfo.ForeignKeys) > 0 {
		needsUUID = true
	}

	// Write imports section
	fmt.Fprintf(writer, "import (\n")
	if needsTime {
		fmt.Fprintf(writer, "\t\"time\"\n")
	}
	if needsJSON {
		fmt.Fprintf(writer, "\t\"encoding/json\"\n")
	}
	if needsUUID {
		fmt.Fprintf(writer, "\t\"github.com/google/uuid\"\n")
	}
	if needsGorm {
		fmt.Fprintf(writer, "\t\"gorm.io/gorm\"\n")
	}
	fmt.Fprintf(writer, ")\n\n")

	// Model struct
	fmt.Fprintf(writer, "type %s struct {\n", modelName)

	// Add BaseModel fields directly with tags
	fmt.Fprintf(writer, "\tID        uuid.UUID       `gorm:\"type:uuid;primaryKey;default:uuid_generate_v4()\"`\n")
	fmt.Fprintf(writer, "\tVersion   int             `gorm:\"type:integer;not null;default:0\"`\n")
	fmt.Fprintf(writer, "\tCreatedAt time.Time       `gorm:\"type:timestamp;not null;default:now()\"`\n")
	fmt.Fprintf(writer, "\tCreatedBy *uuid.UUID      `gorm:\"type:uuid\"`\n")
	fmt.Fprintf(writer, "\tUpdatedAt time.Time       `gorm:\"type:timestamp;not null;default:now()\"`\n")
	fmt.Fprintf(writer, "\tUpdatedBy *uuid.UUID      `gorm:\"type:uuid\"`\n")
	fmt.Fprintf(writer, "\tDeletedAt *gorm.DeletedAt `gorm:\"type:timestamp\"`\n\n")

	// Foreign key fields map to track what we've processed
	foreignKeyFields := make(map[string]bool)

	// First add non-foreign key fields
	for i, columnName := range tableInfo.ColumnNames {
		// Skip if the column is a foreign key or already part of BaseModel
		isForeignKey := false
		isBaseModelField := isBaseModelField(columnName)

		for _, fk := range tableInfo.ForeignKeys {
			if fk.ColumnName == columnName {
				isForeignKey = true
				break
			}
		}

		if !isForeignKey && !isBaseModelField {
			isRequired := strings.Contains(tableInfo.ColumnTags[i], "not null")
			isEnum := tableInfo.EnumTypes[tableInfo.ColumnTypes[i]]
			goType := sqlTypeToGoType(tableInfo.ColumnTypes[i], isRequired, isEnum)
			fieldName := snakeToCamel(columnName)
			// Include GORM tag with proper backticks
			fmt.Fprintf(writer, "\t%s %s `%s`\n", fieldName, goType, tableInfo.ColumnTags[i])
		}
	}

	// Add a blank line if there were any regular fields
	hasRegularFields := false
	for _, columnName := range tableInfo.ColumnNames {
		isBaseModelField := isBaseModelField(columnName)
		isForeignKey := false
		for _, fk := range tableInfo.ForeignKeys {
			if fk.ColumnName == columnName {
				isForeignKey = true
				break
			}
		}
		if !isForeignKey && !isBaseModelField {
			hasRegularFields = true
			break
		}
	}

	if hasRegularFields {
		fmt.Fprintf(writer, "\n")
	}

	// Add foreign key ID fields
	hasForeignKeys := false
	for _, fk := range tableInfo.ForeignKeys {
		if foreignKeyFields[fk.ColumnName] {
			continue // Skip if already processed
		}

		// Format ID field properly (e.g., province_id -> ProvinceID)
		fieldName := formatIDField(fk.ColumnName)

		// Generate GORM tag for foreign key with proper backticks
		fmt.Fprintf(writer, "\t%s uuid.UUID `gorm:\"type:uuid;column:%s\"`\n", fieldName, fk.ColumnName)
		foreignKeyFields[fk.ColumnName] = true
		hasForeignKeys = true
	}

	// Add a blank line if there were any foreign key fields
	if hasForeignKeys {
		fmt.Fprintf(writer, "\n")
	}

	// Add relationship fields
	for _, fk := range tableInfo.ForeignKeys {
		// Get the proper model name for the referenced table
		refModelName := generateModelName(fk.ReferencedTable)

		// Generate GORM tag for relationship with proper backticks
		fmt.Fprintf(writer, "\t%s *%s `gorm:\"foreignKey:%s;references:%s\"`\n",
			refModelName, refModelName, formatIDField(fk.ColumnName), "ID")
	}

	// Check for reverse relationships (where this table is referenced by others)
	for otherTableName, otherTableInfo := range allTableInfos {
		if otherTableName == tableInfo.Name {
			continue // Skip self
		}

		// Look for foreign keys in other tables that reference this table
		for _, fk := range otherTableInfo.ForeignKeys {
			if fk.ReferencedTable == tableInfo.Name {
				// Get the proper model name for the other table
				otherModelName := generateModelName(otherTableName)

				// Use a properly pluralized relationship name
				relationshipFieldName := pluralize(otherModelName)

				// Generate GORM tag for reverse relationship with proper backticks
				otherFKFieldName := formatIDField(fk.ColumnName)
				fmt.Fprintf(writer, "\t%s []%s `gorm:\"foreignKey:%s\"`\n",
					relationshipFieldName, otherModelName, otherFKFieldName)
			}
		}
	}

	fmt.Fprintf(writer, "}\n\n")

	// Add BeforeUpdate method
	fmt.Fprintf(writer, "func (r *%s) BeforeUpdate(tx *gorm.DB) (err error) {\n", modelName)
	fmt.Fprintf(writer, "\tr.Version++\n")
	fmt.Fprintf(writer, "\treturn\n")
	fmt.Fprintf(writer, "}\n")

	// Flush the writer
	return writer.Flush()
}

// Helper function to check if a column is part of the BaseModel
func isBaseModelField(columnName string) bool {
	baseModelFields := map[string]bool{
		"id":         true,
		"version":    true,
		"created_at": true,
		"created_by": true,
		"updated_at": true,
		"updated_by": true,
		"deleted_at": true,
	}

	return baseModelFields[columnName]
}

// GenerateModels generates Go model files from SQL files
func (c *SQLToModelConverter) GenerateModels() error {
	// Get all table infos
	allTableInfos, err := c.getAllTableInfos()
	if err != nil {
		return err
	}

	// Create destination directory if it doesn't exist
	if _, err := os.Stat(c.DestinationPath); os.IsNotExist(err) {
		if err := os.MkdirAll(c.DestinationPath, 0755); err != nil {
			return err
		}
	}

	// Process each table
	for _, tableInfo := range allTableInfos {
		if err := c.generateModelFile(tableInfo, allTableInfos); err != nil {
			fmt.Printf("Error generating model for %s: %v\n", tableInfo.Name, err)
			continue
		}
	}

	return nil
}
