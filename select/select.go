package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	/*
		http.Get(str) performs HTTP GET request against the URL
	*/

	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measuerResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
