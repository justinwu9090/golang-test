package main

import (
	"context"
	"fmt"
	"net/http"
)

// ======================================================
// Interfaces
// ======================================================

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

// ======================================================
// Functions
// ======================================================

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return // todo: log error however you like
		}
		fmt.Fprint(w, data)
	}
}
