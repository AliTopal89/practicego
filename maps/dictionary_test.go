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
        _, got := dictionary.Search("unknown")
        
        assertError(t, got, ErrNotFound)
    })
}

func TestAdd(t *testing.T){
	dictionary := Dictionary{}
	word := "test"
	definition := "this is another test"
	dictionary.Add(word, definition)

	assertDefinition(t, dictionary, word, definition)

}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string){
	t.Helper()

    got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}
    t.Logf("Found the added word")
	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want{
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	
	if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}