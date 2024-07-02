package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")	
	fmt.Println(match)
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// The Submatch variants include information about both the whole-pattern matches and the submatches within 
	// those matches. For example this will return information for both p([a-z]+)ch and ([a-z]+).
	fmt.Println("string submatch:", r.FindStringSubmatch("peach punch"))

	fmt.Println("find string submatch index:", r.FindStringSubmatchIndex("peach punch"))

	/// s string: The input string to search.
	//  n int: The maximum number of matches to return.
	//  -1: the mathod returns all matches 
	// if 0: the method return 0 matches.
	// if >0: the method returns n matches.
	fmt.Println("Find all string matches: ", r.FindAllString("peach punch pinch", -1))

	fmt.Println("all: ", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println("Find 2 matches: ", r.FindAllString("peach punch pinch", 2))
	fmt.Println("Match using []byte instead of string: ", r.Match([]byte("peach")))
	
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp: ", r)
	fmt.Println("replace the string \"a peach\" to: ", r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}