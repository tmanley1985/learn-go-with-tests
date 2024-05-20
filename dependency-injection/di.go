package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Why am I calling this concreteGreet? Because it's asking for a pointer
// to a bytes.Buffer but what we really need to do is to code to an interface
// and the general interface that bytes.Buffer is using is an io.Writer interface!
// Coming from OOP, this concept of writing to an interface is very familiar to me.
func ConcreteGreet(writer *bytes.Buffer, name string) {

	// Fprintf is like Printf except that it takes something
	// that implements the io.Writer interface. Well, bytes.Buffer
	// just so happens to implement that interface!
	//
	// Printf is a convenient wrapper around Fprintf that defaults to stdout.
	// What this allows us to do is accept a writer which is an interface to
	// "put stuff somewhere" and instead of putting it in stdout, we're putting
	// it in a buffer *in our test* so that we can actually test it.
	// Capturing stdout is tricky, so we're injecting this dependency.
	fmt.Fprintf(writer, "Hello, %s", name)
}

// Notice how this accepts an interface? It's not expecting a pointer necessarily
// just anything that adheres to the io.Writer interface.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// Woah! So since we're coding to an interface, our greet function is able to
// actually write to a client on the internet as well! the http.ResponseWriter
// implements the io.Writer interface.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// If you run this file: `go run .` then you'll see that on localhost:5001
// you'll see Hello World!
func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}