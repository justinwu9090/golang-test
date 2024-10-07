package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// goal:
// print 3
// print 3,2,1 and go!
// wait a second between each line
// spys are a kind of "mock" which can record how a dependency is used. They can record the arguments sent in, how many times it's called, etc.. in our case, we keep track how many times Sleep() is called so we can check it in our test.

/*
The advantages of mocking are:
	test that code pauses
	calling a service that will fail  (??)
	test your system in a particular state

it will really test things w/o having to set up databases and 3rd party things

we have covered spies which are a kind of mock.
mocks are actually a type of test double - a generic term where u replace production object for testing purposes
other test doubles are stubs, spies, mocks, ...

https://martinfowler.com/bliki/TestDouble.html


*/

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

	t.Run("sleep before every p rint", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
