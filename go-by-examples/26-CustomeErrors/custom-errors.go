package main

type argError struct {
	arg int
	message string
}

func (e *argError) Error() string {
	return e.message
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
	}
	return arg + 3, nil
}