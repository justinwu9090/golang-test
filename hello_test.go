// run testing on
package main

import "testing" // a testing module i guess

// writing tests needs:
// - file xxx_test.go
// - only takes 1 argumen t *Testing.T
// - need to import "testing"
// - t is of type *testing.T???
func TestHello(t *testing.T) { // don't know what t is here, pointer?
	got := Hello("World") // := is declaring a variable
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want) // calling Errorf method ON our t
	}
}
