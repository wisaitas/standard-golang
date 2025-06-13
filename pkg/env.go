package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func ReadConfig(param any, configPath ...string) error {
	if param == nil {
		return fmt.Errorf("[Share Package Config] : param is nil")
	}

	val := reflect.ValueOf(param)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("[Share Package Config] : param must be a struct")
	}

	path := "."
	if len(configPath) > 0 && configPath[0] != "" {
		path = configPath[0]
	}

	return processStruct(val, "", "", path)
}

func processStruct(val reflect.Value, viperPrefix string, envPrefix string, configPath string) error {
	typ := val.Type()

	configViper := viper.New()
	hasConfigFile := false

	configFiles := []string{
		filepath.Join(configPath, "env.yml"),
		filepath.Join(configPath, "env.yaml"),
		filepath.Join(configPath, ".env"),
	}

	for _, configFile := range configFiles {
		if _, err := os.Stat(configFile); err == nil {
			if strings.HasSuffix(configFile, ".env") {
				if err := godotenv.Load(configFile); err == nil {
					hasConfigFile = true
					break
				}
			} else {
				configViper.SetConfigFile(configFile)
				if err := configViper.ReadInConfig(); err == nil {
					hasConfigFile = true
					break
				}
			}
		}
	}

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		tagValue := typ.Field(i).Tag.Get("defaultValue")
		field := val.Field(i)

		snakeFieldName := ToSnakeCase(fieldName)
		viperKey := snakeFieldName

		if viperPrefix != "" {
			viperKey = viperPrefix + "." + snakeFieldName
		}

		fieldEnvName := ToScreamingSnakeCase(fieldName)
		envKey := fieldEnvName

		if envPrefix != "" {
			envKey = envPrefix + "_" + fieldEnvName
		}

		if field.Kind() == reflect.Struct {
			newViperPrefix := viperKey
			newEnvPrefix := envKey

			if err := processStruct(field, newViperPrefix, newEnvPrefix, configPath); err != nil {
				return fmt.Errorf("[Share Package Config] : %w", err)
			}
			continue
		}

		envValue := os.Getenv(envKey)
		if envValue != "" {
			if err := setFieldValue(field, envValue); err != nil {
				return fmt.Errorf("[Share Package Config] : %w", err)
			}
			continue
		}

		if hasConfigFile && configViper.IsSet(viperKey) {
			var configValue string
			switch field.Kind() {
			case reflect.String:
				configValue = configViper.GetString(viperKey)
			case reflect.Int:
				configValue = fmt.Sprintf("%d", configViper.GetInt(viperKey))
			case reflect.Int64:
				configValue = fmt.Sprintf("%d", configViper.GetInt64(viperKey))
			case reflect.Bool:
				configValue = fmt.Sprintf("%t", configViper.GetBool(viperKey))
			default:
				configValue = configViper.GetString(viperKey)
			}

			if err := setFieldValue(field, configValue); err != nil {
				return fmt.Errorf("[Share Package Config] : %w", err)
			}
			continue
		}

		if tagValue != "" {
			if err := setFieldValue(field, tagValue); err != nil {
				return fmt.Errorf("[Share Package Config] : %w", err)
			}
		}
	}

	return nil
}

func setFieldValue(field reflect.Value, value string) error {
	if !field.CanSet() {
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int:
		num, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("invalid int value: %s", value)
		}
		field.SetInt(int64(num))
	case reflect.Bool:
		switch strings.ToLower(value) {
		case "true", "1", "yes", "on":
			field.SetBool(true)
		case "false", "0", "no", "off":
			field.SetBool(false)
		default:
			return fmt.Errorf("invalid bool value: %s", value)
		}
	case reflect.Int64:
		if duration, err := time.ParseDuration(value); err == nil {
			field.SetInt(int64(duration))
		} else {
			num, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid int64 value: %s", value)
			}
			field.SetInt(num)
		}
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}

	return nil
}
