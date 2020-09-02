package search2

import "errors"

// maps can be modified without passing them as a pointer => because map is a reference type
// A nil map behaves like an empty map, but attempts to write to a nil map will cause a runtime panic.
// Never initialize an empty map variable: (ex: var m map[string]string;  it should be: var dictionary = map[string]string{})
type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

// Search find a word in the dictionary.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Version 4: Add inserts a word and definition into the dictionary.
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// Version 5: Add inserts a word, only if it does not exists.
func (d Dictionary) AddNoDups(word, definition string) error {
	_, err := d.Search(word)
	// Switch statement to match on the error and provides an safety net, in case Search returns an error other than ErrNotFound.
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
