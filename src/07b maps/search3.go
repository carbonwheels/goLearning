package search3

// maps can be modified without passing them as a pointer => because map is a reference type
// A nil map behaves like an empty map, but attempts to write to a nil map will cause a runtime panic.
// Never initialize an empty map variable: (ex: var m map[string]string;  it should be: var dictionary = map[string]string{})
type Dictionary map[string]string

// Version 6: Made errors constant; this required us to create our own DictionaryErr type which implements the error interface.
const (
	// ErrNotFound means the definition could not be found for the given word
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	// ErrWordExists means you are trying to add a word that is already known
	ErrWordExists = DictionaryErr("cannot add word because it already exists")

	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr are errors that can happen when interacting with the dictionary.
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

// Version 7: Update changes the definition of a given word.
func (d Dictionary) Update(word, definition string) error {
	d[word] = definition
	return nil
}
