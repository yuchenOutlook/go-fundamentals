package main

import (
	"fmt"
	s "strings"
)

// We alias fmt.Println to a shorter name as we’ll use it a lot below.
var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "es"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"mother", "father"}, " and "))
	p("Repeat: ", s.Repeat("a", 5))
	p("Replace: ", s.Replace("fooffefefef", "f", "bd", -1))
	p("Replace: ", s.Replace("fooffefefef", "f", "bd", 2))
	p("Split: ", s.Split("a-b-c-d-e", "-"))
	p("ToLower: ", s.ToLower("TEST"))
	p("ToUpper: ", s.ToUpper("test"))
}