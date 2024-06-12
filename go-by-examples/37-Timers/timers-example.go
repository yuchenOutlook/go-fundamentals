package main

import (
	"fmt"
	"time"
)

func main() {
	
	fmt.Println("Main function started.")
	timer1 := time.NewTimer(2 * time.Second)
	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer expired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// you can stop a timer before it expires
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

	/*

The first timer will fire ~2s after we start the program, but the second should be stopped before it has a chance to fire.
	*/
}