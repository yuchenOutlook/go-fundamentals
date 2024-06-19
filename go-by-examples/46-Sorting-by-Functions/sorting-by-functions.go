package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	lenCmp := func (a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age int
	}

	people := []Person {
		{name: "Bob", age: 31},
		{name: "TJ", age: 25},
		{name: "Alex", age: 40},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println("people sorted by age:", people)
}