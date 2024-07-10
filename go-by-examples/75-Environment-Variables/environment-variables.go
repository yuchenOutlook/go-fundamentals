package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("Foo: ", os.Getenv("FOO"))
	fmt.Println("BAR: ", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		fmt.Println(e)
		// pair := strings.SplitN(e, "=", 2)
		// fmt.Println(pair)
	}
}