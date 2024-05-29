package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
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
	ints := intSeq()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println("=====================================init2=====================================")
		
	init2 := intSeq2([]int{1, 2, 3, 4, 5})
	fmt.Println("init2 is created. Now calling it.")
	fmt.Println(init2())
	fmt.Println("the init2 is now done.")
}