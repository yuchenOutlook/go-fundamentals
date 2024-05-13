package main

import "fmt"

type rect struct {
	width, height int
}

// The area method has a receiver type of *rect.
// it takes in the pointer to a rect struct and returns the area.
func (r *rect) area() int {
	return r.width * r.height
}

// The perim method has a receiver type of rect.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height:5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and 
	// pointers for method calls. You may want to use a pointer 
	// receiver type to avoid copying on method calls or to allow 
	// the method to mutate the receiving struct.
	rp := &r
	fmt.Println("rp is the address of r: &r = ", &r)
	fmt.Println("rp's area:", rp.area())
	fmt.Println("rp's perim:", rp.perim())

	rx := r
	fmt.Println("rx is same as r:", rx)
	fmt.Println("rx's area:", rx.area())
	fmt.Println("rx's perim:", rx.perim())
}