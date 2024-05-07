// Slices are an important data type in Go, giving a more powerful
// interface to sequences than arrays.
package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	var s []string
	fmt.Println("Uninit", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("after append, s: ", s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("is len(c) == len(s)? ", len(c) == len(s))

	l := s[2:5]
	fmt.Println(" l = s[2:5]:", l)

	l = s[:5]
	fmt.Println(" l = s[:5]:", l)

	l = s[2:]
	fmt.Println(" l = s[2:]:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declare t:", t)

	t2 := []string{"j", "k", "l"}
	fmt.Println("t is of type:", reflect.TypeOf(t).Kind(), "t2 is of type:", reflect.TypeOf(t2).Kind())
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	} else {
		fmt.Println("t and t2 are not equal")
	}

	// Slices can be composed into multi-dimensional 
	// data structures. The length of the inner slices can vary,
	// unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j:=0; j<innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D slice:", twoD)

}