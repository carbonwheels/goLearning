package search

import (
	"testing"
)

// Version 1: simple search
func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	t.Run("Easy search object", func(t *testing.T) {
		got := Search1(dictionary, "test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
}

// Version 2:
func TestSearch2(t *testing.T) {
	dictionary2 := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary2.Search2("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		//// Version 3: replace assertStrings to assertError so that error message would be constant
		// _, err := dictionary2.Search2("unknown")
		// want := "could not find the word you were looking for"

		// if err == nil {
		// 	t.Fatal("expected to get an error.")
		// }

		// assertStrings(t, err.Error(), want)

		_, got := dictionary2.Search2("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
