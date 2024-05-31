package main

import (
	"errors"
	"fmt"
)

func main() {
	originalError := errors.New("original Error")
	wrappedError := fmt.Errorf("wrapped error layer1: %w", originalError)
	wrappedError_layer2 := fmt.Errorf("wrapped error layer2: %w", wrappedError)
	fmt.Println("wrappedError_layer1: ", wrappedError) // output: wrapped error: Original Error
	fmt.Println("wrappedError_layer2: ", wrappedError_layer2) // output: wrapped error layer2: wrapped error: Original Error

	if unwrappedError := errors.Unwrap(wrappedError_layer2); unwrappedError != nil {
		fmt.Println(unwrappedError) // output: wrapped error layer1: Original Error
	}

	// multiple %w operands
	err1 := errors.New("first error")
    err2 := errors.New("second error")
    multiple_wrapped_errors := fmt.Errorf("multiple errors: %w and %w", err1, err2)
	fmt.Println("wrapped multiple errors:", multiple_wrapped_errors) // output: multiple errors: first error and second error
	if unwrapped_multiple_errors := errors.Unwrap(multiple_wrapped_errors); unwrapped_multiple_errors != nil {
		fmt.Println("unwrapped multiple errors: ", unwrapped_multiple_errors)
	}else {
		fmt.Println("no more wrapped errors")
	}

	joinedError := errors.Join(err1, err2)
	fmt.Println("joined errors: ", joinedError) // two errors are now joined, but returned with new line between each error.
	unwrapped_joined_errors := errors.Unwrap(joinedError)
	fmt.Println("unwrapped joined errors: ", unwrapped_joined_errors) // output: nil
	// https://lukas.zapletalovi.com/posts/2022/wrapping-multiple-errors/
	// see explainations in the post
}