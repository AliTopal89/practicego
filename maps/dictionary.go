package main

const (
	ErrorNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrorWordDoesNotExist = DictionaryErr("cannoy update because the word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

/* You can redirect the user when ErrNotFound is encountered,
   but display an error message when ErrWordDoesNotExist is encountered.
*/
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)

}
