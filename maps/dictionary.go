package main

import (
	"errors"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string	

func (d Dictionary) Search(word string) (string, error) {
	// The `Dictionary` type is the receiver of the `Search` method
	//using an interesting property of the map lookup. It can return 2 values.
	// The second value is a boolean which indicates `if` the key was found successfully.
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}