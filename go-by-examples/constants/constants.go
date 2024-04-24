package main

import (
	"fmt"
	"math"
)


const s string = "constant"
func main() {
	fmt.Println(s)
	const n = 500000000
	const another_n = 50_000_000
	const d = 3e20 / n
	fmt.Println("n value:", n)
	fmt.Println("another_n value:", another_n)
	fmt.Println("true d value:", d)
	fmt.Println("int64 d value:", int64(d))
	fmt.Println("math.Sin(n) value:", math.Sin(n))
}