package jsonscrubber

import (
  "reflect"
)

func AddOnly(s interface{}, fields ...string) interface{} {
  fs := fieldSet(fields...)
	rv := reflect.Indirect(reflect.ValueOf(s))
	out := make(map[string]interface{}, rv.NumField())
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		jsonKey := field.Tag.Get("json")
		if fs[jsonKey] {
			out[jsonKey] = rv.Field(i).Interface()
		}
	}
	return out
}

func fieldSet(fields ...string) map[string]bool {
	set := make(map[string]bool, len(fields))
	for _, s := range fields {
		set[s] = true
	}
	return set
}
