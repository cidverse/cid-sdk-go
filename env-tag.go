package cidsdk

import (
	"os"
	"reflect"
	"strconv"
)

const envVarTag = "env"

// OverwriteFromEnv will overwrite values with the given env values if present
func OverwriteFromEnv(data interface{}) {
	val := reflect.ValueOf(data).Elem()
	t := val.Type()

	// check if the type passed in is a struct
	if t.Kind() != reflect.Struct {
		return
	}

	// iterate over all fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get(envVarTag)
		if tag == "" {
			continue
		}
		if envVal, isSet := os.LookupEnv(tag); isSet {
			fieldVal := val.Field(i)
			switch fieldVal.Kind() {
			case reflect.String:
				fieldVal.SetString(envVal)
			case reflect.Int:
				valAsInt, _ := strconv.Atoi(envVal)
				fieldVal.Set(reflect.ValueOf(valAsInt))
			case reflect.Bool:
				valAsBool, _ := strconv.ParseBool(envVal)
				fieldVal.Set(reflect.ValueOf(valAsBool))
			default:
				// unsupported type
			}
		}
	}
}
