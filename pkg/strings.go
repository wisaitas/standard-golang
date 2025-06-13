package pkg

import (
	"strings"
)

func ToSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

func ToScreamingSnakeCase(str string) string {
	return strings.ToUpper(ToSnakeCase(str))
}
