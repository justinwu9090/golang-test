/*
Goal: write a fucntion which takes a struct and calls the function for all string fields found inside. Useful for JSON parsing

reflection is defined as inspecting the type at runtime and seeing if it matches
*/
package main

// interface{} loses the type safety but now can use "any" type that you don't know at compile time.

func walk(x interface{}, fn func(string)) {
	fn("I still can't believe South Korea beat Germany 2-0 to put them in last in their group")
}
