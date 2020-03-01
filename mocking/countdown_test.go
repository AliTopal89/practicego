package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Logf("testing countdown from 3")
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
