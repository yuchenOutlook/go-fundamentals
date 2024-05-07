package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int) // the map is of type string-> int , string is the key and int is the value
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

    v3 := m["k3"]
    fmt.Println("v3:", v3)

	actual_value, value_exists := m["k2"]
	fmt.Println("actual_value:", actual_value)
	fmt.Println("value_exists:", value_exists)

	for k, v := range m {
		fmt.Println("the k, v in map:", k, v)
		delete(m, k)
	}
	
	// to clear a map, you can use the built-in function delete
	fmt.Println("map after clear:", m)


	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map n:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n and n2 are equal")
	} else {
		fmt.Println("n and n2 are not equal")
	}
}