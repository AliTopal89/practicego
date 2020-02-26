package main

import (
	"bytes"
	"testing"
)

// The "buffer" type from the "bytes" package implements the "Writer" interface.
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
