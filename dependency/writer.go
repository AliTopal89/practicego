package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet sends a personalised greeting to writer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

/*
  `fmt.Fprintf` is like `fmt.Printf` but instead
  takes a `Writer` to send the string to, whereas
  `fmt.Printf` defaults to stdout.
*/

// MyGreeterHandler says Hello, world over HTTP
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))

	if err != nil {
		fmt.Println(err)
	}
}
