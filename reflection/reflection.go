/*
Goal: write a fucntion which takes a struct and calls the function for all string fields found inside. Useful for JSON parsing

reflection is defined as inspecting the type at runtime and seeing if it matches
*/
package main

import "reflect"

// interface{} loses the type safety but now can use "any" type that you don't know at compile time.

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
