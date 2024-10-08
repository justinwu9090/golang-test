/*
Goal: write a fucntion which takes a struct and calls the function for all string fields found inside. Useful for JSON parsing

reflection is defined as inspecting the type at runtime and seeing if it matches
*/
package main

import "reflect"

// interface{} loses the type safety but now can use "any" type that you don't know at compile time.

func walk(x interface{}, fn func(string)) {
	val := getValue(x)
	numberOfValues := 0
	var getField func(int) reflect.Value // stores func pointer either Field() or Index() which returns type reflect.Value
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn) //
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
