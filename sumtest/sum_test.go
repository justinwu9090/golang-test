package main

import (
	"reflect"
	"testing"
)

// arrays have a fixed capacity which is defined at declaration, e.g. //numbers := [5]int{1, 2, 3, 4, 5}
// slices do not encode the size of the collection and thus can have any size
func TestSum(t *testing.T) {

	// t.Run("collection of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}
	// 	got := Sum(numbers)
	// 	want := 15
	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })

	t.Run("collection of any size,", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// want := "bob" --> this has no compile-time checker

	// cannot use != operator for slices. instead use DeepEqual
	// note reflect.DeepEqual is not type-safe.
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
