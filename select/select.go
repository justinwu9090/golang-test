package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	/*
		http.Get(str) performs HTTP GET request against the URL
	*/
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
}
