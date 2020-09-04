package main

import (
	"bytes"
	"testing"
)

// Version 1: test countdown
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Version 2: need to mock sleep data
// Spies are a kind of mock which can record how a dependency is used.
type SpySleeper struct {
	Calls int
}

// They can record the arguments sent in, how many times it has been called, etc.
// In our case, we're keeping track of how many times Sleep() is called so we can check it in our test.
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown2(t *testing.T) {
	buffer := &bytes.Buffer{}
	// We have a dependency on Sleeping which we need to extract so we can then control it in our tests.
	spySleeper := &SpySleeper{}

	Countdown2(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
