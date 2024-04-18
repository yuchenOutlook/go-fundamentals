package main

import (
	"myapp/Scope/packageone"
)

var myVar = "Package level variable"

func main() {
	// variables
	// declare a package level variable for the main
	// package named myVar
	

	// declare a block variable for the main function
	// called blockVar
	var blockVar = "Block level variable"

	// decalre the package level variable in the packageone
	// package named PackageVar

	// create an exported function in package one called PrineMe

	// in the main function, print out the values of myVar,
	// blockVar, and PackageVar on one line
	packageone.PrintMe(myVar, blockVar)
}