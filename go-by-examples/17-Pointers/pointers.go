package main

import (
	"fmt"

	"golang.org/x/crypto/ssh/agent"
)


func zeroval(ival int) {
	ival = 0
}

// you need to input the address of the variable to change the value of the variable
func zeroptr(iptr *int) {
	*iptr = 0
}

func squareval(v int) int {
	v *= v
	fmt.Println("v inside squareval:", v)
	// invalid operation: cannot indirect v (variable of type int)
	// fmt.Println("*v inside squareval:", *v)
	fmt.Println("&v inside squareval:", &v)
	return v
}

func squareAddress(v *int) {
	*v *= *v
	fmt.Println("*v inside squareAddress:", *v)
}

func main() {
	j, z := 25, 50
	fmt.Println("j, z:", j, z)   // 25 50 
	fmt.Println("Address of j, z:", &j, &z)
	fmt.Println("+=====================================================================+")

	p := &j

	fmt.Println("p=&j is the pointer to j:", p)
	fmt.Println("*p is the value of j:", *p)
	fmt.Print(" We also call this as dereferencing the pointer\n")
	fmt.Println("+=====================================================================+")

	*p = 21
	fmt.Println("Now we change the value of j by *p = 21")
	fmt.Println("j value is now:", j)
	fmt.Println("+=====================================================================+")
	
	p = &z
	fmt.Println("Now we change the pointer to point to z by p = &z")
	fmt.Println("Now p is:", p)
	fmt.Println("Now *p is:", *p)
	fmt.Println("+=====================================================================+")

	*p = *p / 5
	fmt.Println("Now we change the value of z by *p = *p / 5")
	fmt.Println("z value is now:", z)
	fmt.Println("p=", p, " This is because p is still representing the address of z")
	fmt.Println("+=====================================================================+")


	x := 10
	squareval(x)
	fmt.Println("x after squareval:", x)
	fmt.Println("+=====================================================================+")

	squareAddress(&x)
	fmt.Println("x after squareAddress:", x)
	fmt.Println("The function squareAddress changes the value of x because it is passed by reference")
	fmt.Println("+=====================================================================+")

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
	
}