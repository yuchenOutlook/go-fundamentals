package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			//errors.Is checks that a given error (or any error in its chain) 
			// matches a specific error value. This is especially useful 
			// with wrapped or nested errors, allowing you to identify specific error 
			//types or sentinel errors in a chain of errors.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("we should buy new tea!")
			}else if errors.Is(err, ErrPower) {
				fmt.Println("we should check the power supply!")
			}else{
				fmt.Println("unknown error: %s\n", err)
			}
		}
		continue
	}
	fmt.Println("Tea is ready!")
}