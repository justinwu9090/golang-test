package main

import (
	"bytes"
	"testing"
)

// use case: test the printout from printf but control where the prints go so that you can write tests against it

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{} // Buffer type - implements Writer interface
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
