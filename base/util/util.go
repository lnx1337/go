package util 

import (
	"reflect"
	"fmt"
)

// convert struct to map
func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			if tagv != "templateId" && tagv != "parentId" {
				data := v.Field(i).Interface()
				if fi.Type.Kind() == reflect.Struct && tagv == "detail" {
					data, _ = ToMap(v.Field(i).Interface(), "json")
				}
				out[tagv] = data
			}
		}
	}
	return out, nil
}