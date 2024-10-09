package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		AssertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup //syncing tool
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait() //same as CUDA __syncthreads()

		AssertCounter(t, &counter, wantedCount)
	})
}

func AssertCounter(t testing.TB, got *Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d want %d", got, want)
	}
}

/*
Mutex allows us to add locks to our data
WaitGroup is a means of waiting for goroutines to finish jobs

decision making:
Use channels when passing ownership of data
Use mutexes for managing state
*/
