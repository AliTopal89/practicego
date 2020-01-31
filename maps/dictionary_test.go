package main

import (
	"testing"
)
//A map is declared using the following syntax -
	// var m map[KeyType]ValueType
func TestSearch(t *testing.T) {
	// t of type *testing.T is your "hook" into the testing framework 
	//so you can do things like t.Fail() when you want to fail.
	dictionary := Dictionary{"test": "this is just a test"}
    t.Logf("testing known word")
    t.Run("known word", func(t *testing.T) {
        got, _ := dictionary.Search("test")
        want := "this is just a test"

        assertStrings(t, got, want)
    })
    t.Logf("testing unknown word")
	t.Run("unknown word", func(t *testing.T) {
        _, err := dictionary.Search("unknown")
        want := "could not find the word you were looking for"

        if err == nil {
            t.Fatal("expected to get an error.")
        }

        assertStrings(t, err.Error(), want)
    })
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	
	if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}