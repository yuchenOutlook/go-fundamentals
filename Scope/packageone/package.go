package packageone

import "fmt"

var privateVar = "I am private"
var PublicVar = "I am public"

func notExported() {
	fmt.Println("I am not exported")
}

func Exported() {
	fmt.Println("I am exported")
}