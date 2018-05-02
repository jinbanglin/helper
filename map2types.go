package helper

import (
	"reflect"
	"strings"
)

func Struct2Map(in interface{}, tag string) map[string]string {
	out := make(map[string]string)
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
			out[strings.TrimSuffix(tag, ",omitempty")] = v.Field(i).String()
		}
	}
	return out
}

func Struct2MapString(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).String()
	}
	return data
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
