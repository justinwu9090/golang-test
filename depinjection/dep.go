package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "Elodie\n")
	// example for when to use io writer

	// when writing an HTTP handler, you're given a response writer and request pointer that was used to make the request.
	// when you implement your server you write your response using the writer. so http.ResponseWriter also implements io.Writer.
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
