/*
Goal: write a fucntion which takes a struct and calls the function for all string fields found inside. Useful for JSON parsing

reflection is defined as inspecting the type at runtime and seeing if it matches
*/
package main

import "reflect"

// interface{} loses the type safety but now can use "any" type that you don't know at compile time.

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		// this is a while loop
		for {
			// if condition checks ok which is extracted out of the channel
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		// calling function w/ arbitrary number of return values
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
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
