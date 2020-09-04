package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

//  Version 1: execution for the defined Duration.
func Countdown(out io.Writer) {
	// Countdown prints a countdown from 3 to out.
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	time.Sleep(4 * time.Second)
	// Upon completion display the final word
	fmt.Fprint(out, finalWord)
}

//  Version 2: Sleep will pause execution for the defined Duration.
// This lets us then use a real Sleeper in main and a spy sleeper in our tests.
// By using an interface our Countdown function is oblivious to this and adds some flexibility for the caller.

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is an implementation of Sleeper with a predefined delay.
type DefaultSleeper struct{}

// Sleep will pause execution for the defined Duration.
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper.
func Countdown2(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)

	sleeper := &DefaultSleeper{}
	Countdown2(os.Stdout, sleeper)
}
