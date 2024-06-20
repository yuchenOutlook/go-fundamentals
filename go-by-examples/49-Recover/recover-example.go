package main

import (
	"fmt"
	"time"
)

func myPanic() {
	fmt.Println("Starting myPanic, wait for 2 seconds before panicking...")
	time.Sleep(2 * time.Second)
	panic("a problem occurred")
}

func main() {
	
	defer func() {
		// recover()
		// fmt.Println("Recovered from panic")
		// fmt.Println("Can I recover twice?")
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	myPanic()
	fmt.Println("Program continues after panic")
}