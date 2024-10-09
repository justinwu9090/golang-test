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
		store.Cancel()
		fmt.Fprint(w, store.Fetch())
	}
}
