package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// In Go, when a variable's type is declared, the type of that
	// variable is indeed immutable—meaning the variable's type cannot be 
	// changed after its declaration. However, 
	// it's important to note that the immutability of the type does 
	// not mean that the value of the variable itself is immutable.

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}