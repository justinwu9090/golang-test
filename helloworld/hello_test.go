// run testing on
package main

import "testing" // a testing module i guess

// writing tests needs:
// - file xxx_test.go
// - only takes 1 argumen t *Testing.T
// - need to import "testing"
// - t is of type *testing.T???
func TestHello(t *testing.T) { // don't know what t is here, pointer?
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("World", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() // tells the test suite that this is a helper, so the failing line number doesn't end up here
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
