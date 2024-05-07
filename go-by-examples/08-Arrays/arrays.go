package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	fmt.Println("length of a is:", len(a))

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get a[4]:", a[4])

	b := [5]int {1, 2, 3, 4, 5}
	c := [5]float64 {1.1, 2.2, 3.3, 4.4, 5.59}
	fmt.Println("declare b:", b)
	fmt.Println("declare c:", c)

	// You can also have the compiler count the number of elements for you with ...
	b = [...]int{1, 2, 3, 4, 5}
	d := [...]bool{true, false, true, false}
	fmt.Println("another way to declare b with 3 dots [...]int{1, 2, 3, 4, 5}:", b)
	fmt.Println("another way to declare d with 3 dots [...]bool{true, false, true, false}:", d)


	b = [...]int{100, 3:400, 500}
	fmt.Println("another way to declare b with 3 dots [...]int{100, 3:400, 500}:", b)

	// 2D arrays
	var twoD [2][3]int
	for i:=0; i< 2; i++ {
		for j:=0; j<3; j++ {
			twoD[i][j] = i+j
		}
	}

	fmt.Println("2D array:", twoD)

	twoD = [2][3]int {
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2D array directly declare:", twoD)

}

