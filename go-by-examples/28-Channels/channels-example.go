package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Waiting 2 seconds before message being sent")
		messages <- "ping"
	} ()

	fmt.Println("Waiting for message")
	msg := <- messages
	fmt.Println("Message received")
	fmt.Println(msg)
}