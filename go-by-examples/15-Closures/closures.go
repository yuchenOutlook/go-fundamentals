package main

import (
	"fmt"
	"math/rand"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

		
func randomfunc() func() int {
	p := fmt.Println
	p("randomfunc is called.")
	v := rand.Intn(100)
	return func() int {
		p("the function within the randomfunc is called, now returning a random int: ", v)
		return v;
	}
}

func intSeq2(slices []int) func() int {
	for i := range slices {
		slices[i] *= 2
	}

	return func() int {
		sum := 0
		for _, v := range slices {
			sum += v
		}
		fmt.Println("The annoynimous function is executed.")
		fmt.Println("sum:", sum)
		return sum
	}
}

func main() {
	p := fmt.Println
	r := randomfunc
	print(r()())
	ints := intSeq()
	p(ints())
	p(ints())
	p(ints())

	newInts := intSeq()
	p(newInts())
	p("=====================================init2=====================================")
		
	init2 := intSeq2([]int{1, 2, 3, 4, 5})
	p("init2 is created. Now calling it.")
	p(init2())
	p("the init2 is now done.")
}