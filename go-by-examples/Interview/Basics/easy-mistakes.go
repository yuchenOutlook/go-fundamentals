package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	for _, value := range mySlice {
		value *= 2
	}

	fmt.Println(mySlice) // This result will still be [1, 2, 3, 4, 5]
	// This is because the value variable is a copy of the slice element, not a reference to the element itself

	for i, _ := range mySlice {
		mySlice[i] *= 2
	}

	fmt.Printf("values are %+v\n", mySlice) // This result will be [2, 4, 6, 8, 10]
}