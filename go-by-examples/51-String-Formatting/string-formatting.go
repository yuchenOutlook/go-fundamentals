package main

import (
	"fmt"
	"os"
)

// Go offers excellent support for string formatting in the printf tradition. Here are some examples of common string formatting tasks.
type point struct {
	x, y int
}
var pf = fmt.Printf
func main() {
	p := point{1, 2}

	// Just print the value of the struct
	fmt.Printf("struct1: %v\n", p)

	// print the struct field names with value
	fmt.Printf("struct2: %+v\n", p)

	// print the struct field names with value and type
	fmt.Printf("struct3: %#v\n", p)

	fmt.Printf("Just the Type: %T\n", p)

	fmt.Printf("print the bool: %t\n", false)

	fmt.Printf("print the decimal: %d\n", 12345)
	fmt.Printf("print the float: %f\n", 12345.34343435345)
	fmt.Printf("Print binary representation: %b\n", 14)
	fmt.Printf("print character corresponding to the integer (for 33) : %c\n", 33)
	fmt.Printf("print the hex encoding (for 456): %x\n", 456)
	fmt.Printf("float2 : %e\n", 12345.34343435345)
	fmt.Printf("float3 : %E\n", 12345.34343435345)
	fmt.Printf("basic string printing %s\n", "\"string\"")
	fmt.Printf("Double quoted string: %q\n", "\"string\"")
	pf("Hex this string: %x\n", "hex this")
	pf("pointer: %p\n", &p)
	pf("Control digit width: |%6d|%6d|\n", 12, 345)
	pf("Width2 with floating point number: |%6.4f|%6.4f|\n", 1.2, 3.45)
	pf("Left justify: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	pf("control witdth when formatting strings in table like output: |%6s|%6s|\n", "foo", "b")
	pf("Control length when formatting strings: |%.5s|%.5s|\n", "abcdef", "abcdef")
	pf("left justify strings: |%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats and returns a strign without print it anywhere
	s := fmt.Sprintf("a %s\n", "string")
	fmt.Println(s)

	// Fprintf formats and writes to io.Writers
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}