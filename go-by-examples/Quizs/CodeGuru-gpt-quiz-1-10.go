package main

import (
	"fmt"
	"reflect"
)

// Quiz on Go Programming: Fill in the blanks or answer the questions in comments.

// Question 1: What is printed when you run the hello-world.go program? (Answer with a comment)
func question1() {
	// fmt.Println("Hello, World!")
	// Hello, World!
}

// Question 2: Write code to declare an integer variable and initialize it with value 42.
func question2() {
	// var x int = ___
	var x int = 42
	fmt.Println(x)
}

// Question 3: Modify the loop in for.go to count from 1 to 10 inclusive.
func question3() {
	// for i := 0; i < __; i++ {
	// 	 fmt.Println(i)
	// }
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// Question 4: Using the switch.go file as a reference, write a switch that prints "small" if i < 10, "big" if i >= 10.
func question4(i int) {
	// switch {
	// case ___:
	//	 fmt.Println("small")
	// case ___:
	//	 fmt.Println("big")
	// }

	switch {
		case i < 10:
			fmt.Println("small") 
		case i >= 10:
			fmt.Println("big")
		default:
			fmt.Println("default")
	}
}

// Question 5: From the maps.go file, how do you delete an entry from a map? (Answer with a comment)
func question5() {
	// m := map[string]int{"foo": 1, "bar": 2}
	// ___(m, "bar")
	// delete(m, "bar")
}

// Question 6: Arrays vs Slices - From arrays.go and slices.go, explain the main difference. (Answer with a comment)
// answer: the arrays are fixed in size, while the slices are dynamic in size.

// Question 7: Based on values.go, write code that prints the type of a variable that holds the value true.
func question7() {
	// fmt.Println(__(true))
	fmt.Print(reflect.TypeOf(true).Kind())
}

// Question 8: Given a constant definition in constants.go, define a constant named pi with a value of 3.14.
func question8() {
	// const pi = ___
	const pi = 3.14
}

// Question 9: From if-else.go, rewrite the code snippet to use else if to check additional condition: i == 5
func question9(i int) {
	// if i < 10 {
	// 	 fmt.Println("i is less than 10")
	// } else if ___ {
	//	 fmt.Println("i is exactly 5")
	// } else {
	// 	 fmt.Println("i is greater than or equal to 10")
	// }
	if i < 10 {
		 fmt.Println("i is less than 10")
	} else if i == 5 {
		 fmt.Println("i is exactly 5")
	} else {
		 fmt.Println("i is greater than or equal to 10")
	}
}

// Question 10: Write a function that takes a slice of strings and returns the first element, if there are any elements.
func question10(slice []string) string {
	// if len(slice) > 0 {
	//	 return ___
	// }
	// return ""
	if len(slice) > 0 {
		 return slice[0]
	}
	return ""
}