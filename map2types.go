package helper

import (
	"reflect"
)

func Struct2Map(obj interface{}, tag string) map[string]string {
	out := make(map[string]string)
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if tag := fi.Tag.Get(tag); tag != "" {
			out[tag] = v.Field(i).String()
		}
	}
	return out
}

func Struct2MapInterface(in interface{}, tag string) map[string]interface{} {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if tag := fi.Tag.Get(tag); tag != "" {
			out[tag] = v.Field(i).Interface()
		}
	}
	return out
}
