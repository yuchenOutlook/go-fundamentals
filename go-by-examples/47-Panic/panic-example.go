package main

import "os"

func main() {
	// panic stops the program immediately and prints a stack trace
	panic("a problem occurred")

	// Note that unlike some languages which use exceptions for handling of many errors, in Go 
	// it is idiomatic to use error-indicating return values wherever possible.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}