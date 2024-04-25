package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("two")
	default:
		fmt.Print("unknown")
	}
	fmt.Println("")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("It's a weekday")
	default:
		fmt.Println("It's a mystery")
	}

	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("It's before noon")
		default:
			fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
			case bool:
				fmt.Println("I'm a bool")
			case int:
				fmt.Println("I'm an int")
			case string:
				fmt.Println("I'm a string")

			default:
				fmt.Printf("The type is %T\n", t)
				// fmt.Println("Don't know type", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")
	whatAmI(3.14i)
	

}