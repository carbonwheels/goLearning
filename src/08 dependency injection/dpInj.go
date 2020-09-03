package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Version 1: Greet sends a personalized greeting to writer.
func Greet(writer io.Writer, name string) {
	// We could simplify the code to below -
	// However Printf calls Fprintf, we need to do is to be able to inject (aka pass in) the dependency of printing.
	fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}

// version 2: MyGreeterHandler says Hello, world over HTTP.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Version 1: Print out command line
	Greet(os.Stdout, "Elodie")

	// Startup webserver & list to 5000
	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))

	if err != nil {
		fmt.Println(err)
	}
}
