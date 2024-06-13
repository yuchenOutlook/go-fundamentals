package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout after 1 second")
	}

	elapsed1 := time.Since(start)
	fmt.Printf("Time elapsed1: %s\n", elapsed1)

	start = time.Now()
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
		case res := <-c2:
			fmt.Println(res)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout after 3 seconds")
	}
	elapsed2 := time.Since(start)
	fmt.Printf("Time elapsed2: %s\n", elapsed2)
}