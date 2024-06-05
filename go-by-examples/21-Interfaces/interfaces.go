package main

import (
	"fmt"
	my_interface_example "myapp/go-by-examples/21-Interfaces/my-interface-example"
)

// Speaker interface definition
type Speaker interface {
    Speak() string
    Speak2() int
}

// Person struct definition
type Person struct {
    Name string
}

func (pp Person) Speak2() int {
     fmt.Println("This is Speak2 method, person name is:" + pp.Name)
     return 2
}

// Speak method implementation by Person
func (p Person) Speak() string {
    return "Hi, my name is " + p.Name
}

// introduce takes a Speaker and calls its Speak method
func introduce(speaker Speaker) {
    fmt.Println(speaker.Speak())
    fmt.Println(speaker.Speak2())
}

func main() {
    john := Person{Name: "John"}
    introduce(john)  // Output: Hi, my name is John

	javaProgrammer := my_interface_example.JavaProgrammer{
		Name:       "John",
		Age:        30,
		Occupation: "Java Programmer",
	}
    javaProgrammer.Develop() // Output: John is developing a new feature

    golangProgrammer := my_interface_example.NewGolangProgrammer("John", "New York")
    golangProgrammer.Develop() // Output: John is developing a new feature at New York
    golangProgrammer.WriteUnitTests("feature1") // Output: John is writing unit tests for feature1
    golangProgrammer.Debug("feature1") // Output: John is debugging feature1
    

}
