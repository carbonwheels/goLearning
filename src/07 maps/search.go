package search

import "errors"

// Version 1: Search find a word in the dictionary.
func Search1(dictionary map[string]string, word string) string {
	// Getting a value out of a Map is the same as getting a value out of Array map[key].
	return dictionary[word]
}

// Version 2: Dictionary store definitions to words.
// Dictionary type which acts as a thin wrapper around map. With the custom type defined, we can create the Search method.
type Dictionary map[string]string

// Version 3: replace assertStrings to assertError so that error message would be constant
var ErrNotFound = errors.New("could not find the word you were looking for")

// Version2: Set the object with method Search2 and have 2 items returned
func (d Dictionary) Search2(word string) (string, error) {
	// since d is of Dictionary (ex. dictionary[word]) it will return a tuple
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
