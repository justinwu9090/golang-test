/*
Goal: write a fucntion which takes a struct and calls the function for all string fields found inside. Useful for JSON parsing

reflection is defined as inspecting the type at runtime and seeing if it matches
*/
package main

import "reflect"

// interface{} loses the type safety but now can use "any" type that you don't know at compile time.

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x) //reflect package func that returns a Value of a given variable.
	// assumes: 1. look at first and only field
	field := val.Field(0)
	fn(field.String()) // 2. call String() which returns underlying value a sstring
}
