package cidsdk

import (
	"os"
	"reflect"
	"strconv"
)

const envVarTag = "env"

// OverwriteFromEnv will overwrite values with the given env values if present
func OverwriteFromEnv[T any](data *T) {
	t := reflect.TypeOf(*data)
	v := reflect.ValueOf(*data)

	// iterate over all fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get(envVarTag)
		if tag == "" {
			continue
		}

		// overwrite with env value if present
		if tag != "" {
			if field.Type.Kind() == reflect.Int {
				val, isSet := os.LookupEnv(tag)
				if isSet {
					valAsInt, _ := strconv.Atoi(val)
					v.Elem().Field(i).SetInt(int64(valAsInt))
				}
			} else if field.Type.Kind() == reflect.String {
				val, isSet := os.LookupEnv(tag)
				if isSet {
					v.Elem().Field(i).SetString(val)
				}
			} else if field.Type.Kind() == reflect.Bool {
				val, isSet := os.LookupEnv(tag)
				if isSet {
					valAsBool, _ := strconv.ParseBool(val)
					v.Elem().Field(i).SetBool(valAsBool)
				}
			}
		}
	}
}
