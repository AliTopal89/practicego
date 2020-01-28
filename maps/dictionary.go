package main

type Dictionary map[string]string	

func (d Dictionary) Search(word string) string {
	// The `Dictionary` type is the receiver of the `Search` method
	return d[word]
}