package main

import (
	"testing"

)
//A map is declared using the following syntax -
	// var m map[KeyType]ValueType
func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}