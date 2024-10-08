package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	// default tensecondTimeout using Racer
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// "select" will launch every case and only execute the case that returns first.
	// this is a replacement for myVar := <-ch which is blocking for each val that comes out of the channel
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
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

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
