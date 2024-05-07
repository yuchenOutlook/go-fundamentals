package main

import (
	"fmt"
	"time"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	for num := 9; num > 7; num-- {
		fmt.Println(num, "is bigger than 7")
	}

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("two")
	case 3:
		fmt.Print("three")
	default:
		fmt.Print("unknown")
	}
	fmt.Println()
	fmt.Println("Today is", time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("It's a weekday")
	}

	whatTypeAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type", t)
		}
	}
	whatTypeAmI(true)
	whatTypeAmI(1)
	whatTypeAmI("hey")

}
