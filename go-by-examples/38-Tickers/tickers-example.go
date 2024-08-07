package main

import (
	"fmt"
	"time"
)

/*

Timers are for when you want to do something once in the future -
tickers are for when you want to do something repeatedly at regular intervals.
Here’s an example of a ticker that ticks periodically until we stop it.
*/
func main() {
	fmt.Println("Main function started.")
	ticker := time.NewTicker(500 * time.Millisecond)
	fmt.Println("Ticker started")
	done := make(chan bool)

	go func ()  {
		i := 0
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				i++
				fmt.Printf("Tick %d at %v \n", i, t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}