package main

import "fmt"

func main() {
	nums := []int {1, 2, 3, 4, 5}
	sum := 0
	for index, num := range nums {
		sum += num
		if num % 2 == 0 {
			// number is even
			fmt.Print("number is even: ", num)
			fmt.Println("index: ", index, "num: ", num, "sum: ", sum)
		}
		fmt.Print("number is odd: ", num)
		fmt.Println("index: ", index, "num: ", num, "sum: ", sum)
	}
	fmt.Println("sum: ", sum)

	kvs := map[string]string {"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for key := range kvs {
		fmt.Println("key:", key)
	}

	for i, c := range "go" {
		fmt.Println("index in \"go\":", i, "character is \"go\":", c)
	}

}