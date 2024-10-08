package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		AssertValue(t, counter, 3)
	})
}

func AssertValue(t testing.TB, got Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d want %d", got, want)
	}
}
