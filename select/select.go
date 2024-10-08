package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	// "select" will launch every case and only execute the case that returns first.
	// this is a replacement for myVar := <-ch which is blocking for each val that comes out of the channel
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {

	// always "make" channels. var ch chan struct{} is a nil channel and sending to it will do nothing
	ch := make(chan struct{})
	// go routine here
	go func() {
		// http.Get(str) performs HTTP GET request against the URL
		http.Get(url)
		close(ch)
	}()

	// return the channel?
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
