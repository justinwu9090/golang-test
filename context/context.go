package context

import (
	"fmt"
	"net/http"
)

// ======================================================
// Interfaces
// ======================================================

type Store interface {
	Fetch() string
	Cancel()
}

// ======================================================
// Functions
// ======================================================

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		// fetch must be done in a goroutine to race against the context's Done()
		go func() {
			data <- store.Fetch()
		}()

		// use select to race said data vs the Done() which signals "cancelled".
		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
