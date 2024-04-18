package main

import (
	"fmt"
	"myapp/Scope/packageone"
)

var one = "this is a global variable"
func main() {
	var one = "this is a local variable"
	var somethingElse = "this is a block level variable"
	fmt.Println(one)
	fmt.Println(somethingElse)
	fmt.Println(packageone.PublicVar)
	// This will not work
	// fmt.Println(packageone.privateVar)
	myFunc()
	packageone.Exported()
}

func myFunc() {
	fmt.Println(one)
}