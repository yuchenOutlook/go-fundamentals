package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := intSeq()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	newInts := intSeq()
	fmt.Println(newInts())
}