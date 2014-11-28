package helper

import (
	"reflect"
)

func Struct2Map(s interface{}) map[string]interface{} {

	v := reflect.ValueOf(s)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	m := make(map[string]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}
	return m
}
func StructSlice2MapSlice(s interface{}) []map[string]interface{} {
	sv := reflect.ValueOf(s)

	ret := make([]map[string]interface{}, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		strct := sv.Index(i)

		v := strct
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		t := v.Type()
		m := make(map[string]interface{}, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			m[t.Field(i).Name] = v.Field(i).Interface()
		}
		ret = append(ret, m)
	}
	return ret
}
