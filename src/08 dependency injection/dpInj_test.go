package main

import (
	"bytes"
	"testing"
)

// Version 1: Testing Greet function and passing buffer - abstracting the code testable and more reusable.
func TestGreet(t *testing.T) {
	// The buffer type from the bytes package implements the Writer interface.
	buffer := bytes.Buffer{}
	// Under the covers we are using Writer to send our greeting somewhere.
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
