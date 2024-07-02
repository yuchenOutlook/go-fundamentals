package main

import "fmt"

type Electronics int

const (
	Phone Electronics = iota
	TV
	PC
	Server
	Tablet
)

func (e Electronics) String() string {
    names := map[Electronics]string{
        Phone: "Phone",
        TV: "TV",
        PC: "PC",
        Server: "Server",
		Tablet: "Tablet",
    }
    return names[e]
}


func main() {
	device1 := TV
	fmt.Println("This device is:", device1)

	device2 := Tablet
	fmt.Println("This device is:", device2)
}