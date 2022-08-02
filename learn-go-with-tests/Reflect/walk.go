package Reflect

import (
	"reflect"
)

func walk(x interface{}, fn func(input string))  {
	//val := getValue(x)
	//if val.Kind() == reflect.Slice {
	//	for i := 0; i < val.Len(); i++ {
	//		walk(val.Index(i).Interface(), fn)
	//	}
	//	return
	//}
	//for i := 0; i < val.NumField(); i++ {
	//	field := val.Field(i)
	//	switch field.Kind() {
	//	case reflect.String :
	//		fn(field.String())
	//	case reflect.Struct :
	//		walk(field.Interface(), fn)
	//	}
	//}
	val := getValue(x)

	//numberOfValues := 0
	//var getField func(int) reflect.Value
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		//numberOfValues = val.NumField()
		//getField = val.Field
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		//numberOfValues = val.Len()
		//getField = val.Index
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}
	//for i := 0; i < numberOfValues; i++ {
	//	walk(getField(i).Interface(), fn)
	//}
}
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}