package main

import (
	"bytes"
	"testing"
)

// goal:

// print 3

// print 3,2,1 and go!

// wait a second between each line

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
