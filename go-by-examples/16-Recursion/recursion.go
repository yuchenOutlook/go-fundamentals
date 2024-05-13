package main

import "fmt"

func fact(n int) int{
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("fact(5):", fact(5))
	var fib func(n int) int

	fib = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println("fib(5):", fib(5))
}