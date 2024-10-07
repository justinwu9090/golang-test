package main

import (
	"bytes"
	"testing"
)

// goal:
// print 3
// print 3,2,1 and go!
// wait a second between each line
// spys are a kind of "mock" which can record how a dependency is used. They can record the arguments sent in, how many times it's called, etc.. in our case, we keep track how many times Sleep() is called so we can check it in our test.

// ======================================================
// tests
// ======================================================
func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want %d got %d", countdownStart, spySleeper.Calls)
	}
}
