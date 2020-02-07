package main

import (
	"errors"
)

var ErrorNotFound = errors.New("could not find the word you were looking for")
var ErrorWordExists = errors.New("this word already exists")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// The `Dictionary` type is the receiver of the `Search` method
	//using an interesting property of the map lookup. It can return 2 values.
	// The second value is a boolean which indicates `if` the key was found successfully.
	definition, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil

}
