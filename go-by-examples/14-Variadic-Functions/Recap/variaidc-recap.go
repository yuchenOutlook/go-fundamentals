package main

import "fmt"

func main() {

	fmt.Println(sum(2, 2))
	fmt.Println(sum(2, 3, 53, 5))
	x, y := 2, 5
	fmt.Println(x, y)
	addX, addY := &x, &y
	fmt.Println(sumAdd(addX, addY))
	fmt.Println("x, y:", x, y)

	fmt.Println("=============Another example of variadic function======")
	buy("shoes", 100)
	buy("Seasalt", 100, 10)
}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
		fmt.Println("num:", num)
	}
	return sum
}

func sumAdd(nums ...*int) int {
	sum := 0
	for _, num := range nums {
		sum += *num
		fmt.Println("num:", *num)
	}

	for _, num := range nums {
		*num = 0
	}
	return sum
}

func buy(item string, price int, discount ...int) {
	if len(discount) >= 1 {
		price -= discount[0]
	}
	fmt.Printf("You bought %s at %d\n", item, price)
}