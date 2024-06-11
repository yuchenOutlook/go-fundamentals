package main

import (
	"fmt"
	"time"
)

func main() {
	
	timer1 := time.NewTimer(2 * time.Second)
	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer expired.
	timer1_first, timer1_second := <-timer1.C
	fmt.Println("timer1.C first return value", timer1_first)
	fmt.Println("timer1.C second expired at", timer1_second)
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	// you can stop a timer before it expires
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}