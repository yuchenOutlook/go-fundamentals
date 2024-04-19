package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, just press ENTER when you are ready."

func main() {
	// seed rand number generator
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction
	printTheGame(firstNumber, secondNumber, subtraction, answer)
}

func printTheGame(firstNumber int, secondNumber int, subtraction int, answer int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Guess the Number Game.")
	fmt.Println("=====================")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10 ", prompt)

	reader.ReadString('\n')
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')
	fmt.Println("The answer is", answer)
}