package main

import "fmt"

// Speaker interface definition
type Speaker interface {
    Speak() string
}

// Person struct definition
type Person struct {
    Name string
}

// Speak method implementation by Person
func (p Person) Speak() string {
    return "Hi, my name is " + p.Name
}

// introduce takes a Speaker and calls its Speak method
func introduce(speaker Speaker) {
    fmt.Println(speaker.Speak())
}

func main() {
    john := Person{Name: "John"}
    introduce(john)  // Output: Hi, my name is John
}
