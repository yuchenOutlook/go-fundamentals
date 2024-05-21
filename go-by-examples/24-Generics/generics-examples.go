package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomMap[T comparable, V int | string] map[T]V 

type UserId int

type CustomData interface {
	constraints.Ordered | []byte | []int | []string
}

type User[T CustomData] struct {
	ID int
	Name string
	Data T
}

func Add[T constraints.Ordered] (a, b T) T {
	return a + b
}

func Add2[T ~int] (a, b, c T) T {
	return a + b + c
}

// This mapValues can only take in a slice of integers and a function that takes in an integer and returns an integer
// The function signature is func(int) int
func MapValues(values []int, mapFunc func(int) int ) []int {
	var newValues []int
	for _, value := range values {
		newValue := mapFunc(value)
		newValues = append(newValues, newValue)
	}
	return newValues
}

// Now this mapValuesGeneric can take in a slice of any type that satisfies the constraints.Ordered interface
func MapValuesGeneric[T constraints.Ordered](values [] T, mapFunc func(T) T) []T {
	var newValues []T
	for _, value := range values {
		newValue := mapFunc(value)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	m := make(CustomMap[int, string])
	m[3] = "Hello"
	fmt.Println("This is how you define a custom map data with comparable keys and int or string values:")
	fmt.Println(m)

	result := Add(1, 2)
	fmt.Println(result)

	result2 := Add(1.12, 2.12)
	fmt.Println(result2)

	result3 := Add2(1, 2, 3)
	fmt.Println(result3)

	mapResult := MapValues([]int{1, 2, 3}, func(value int) int {
	{
		return value * 2
	}
	})
	fmt.Println(mapResult)

	// You see this answer is more generic and can be used for any type that satisfies the constraints.Ordered interface
	mapResultGeneric := MapValuesGeneric([]float32{1.2, 2, 3}, func(value float32) float32 {
	{
		return value * 4.2
	}
	})
	fmt.Println(mapResultGeneric)

	u := User[int] {
		ID: 1,
		Name: "John",
		Data: 333,
	}

	u2 := User[string] {
		ID: 2,
		Name : "Jane",
		Data: "This is Jane Data",
	}

	fmt.Println(u)
	fmt.Println(u2)
}