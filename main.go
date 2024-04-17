package main

import (
	"fmt"
	"myapp/doctor"
)

func main() {
	// var whatToSay string = "Hello, World!"
	// oneStep_whatToSay := "Hello, World!"
	// sayHello(whatToSay)
	// sayHello(oneStep_whatToSay)
	var whatToSay string
	whatToSay = doctor.Intro()
	fmt.Println(whatToSay)
}

func sayHello(whatToSay string) {
	fmt.Println(whatToSay)
}