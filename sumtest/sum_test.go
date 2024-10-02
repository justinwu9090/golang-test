package main

import "testing"

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
