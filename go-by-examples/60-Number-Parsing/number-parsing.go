package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("2323", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

	/*
	Simplicity:

	Atoi is simpler to use when you know you are dealing with base 10 integers.
	It does not require you to specify the base and bit size, making the code more concise.
	Common Use Case:

	Atoi is intended specifically for converting strings to integers in base 10, which is a very common requirement.
	*/
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
}
