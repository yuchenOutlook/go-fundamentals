package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <- pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

/*
When you create a channel using `make(chan Type, bufferSize)`, it's a bidirectional channel by default, meaning you can both send to and receive from it. The direction is not specified at the point of creation.

The direction of a channel gets specified when you pass it to a function. In your code, the `ping` function takes a send-only channel (`pings chan<- string`), and the `pong` function takes a receive-only channel (`pings <-chan string`) and a send-only channel (`pongs chan<- string`).

So, in the context of the `ping` function, `pings` is a send-only channel. In the context of the `pong` function, `pings` is a receive-only channel and `pongs` is a send-only channel.

However, in the `main` function, both `pings` and `pongs` are bidirectional because they were created as such and the `main` function doesn't impose any direction restrictions on them.
*/