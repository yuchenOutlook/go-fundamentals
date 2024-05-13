package main

import "fmt"

func subtract(nums ...int) int {
	fmt.Print(nums, "")
	var total int = 0
	for index, num := range nums {
		fmt.Println("index ", index, " is:", num)
		if len(nums) == 0 {
			total = num
		} else {
			total -= num
		}
	}
	return total
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	

	// If you already have multiple args in a slice, 
	// apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4, 5, 6}
	sum(nums...)
	subtract(1, 2, 3, 4, 5, 6)
}