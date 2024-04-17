package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	// write a while loop that will keep running until the user types "quit"

	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	// whenever the user presses enter, the program will read the input and store it in the userInput variable
	for {
		fmt.Println("->")
		userInput, _ := reader.ReadString('\n')

		// This \r\n is for windows, \n is for unix
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}
}