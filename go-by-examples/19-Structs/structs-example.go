package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name, age: 43}
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	k := newPerson("Kio")
	fmt.Println("Created new person kio by calling newPerson(\"kio\"):")
	fmt.Println("k is the address of newPerson:", k)
	fmt.Println("*k is the value of newPerson:", *k)
	k.age = 50
	
	s := person{name: "Sean", age: 32}
	sp := &s // sp is address of s
	fmt.Println("sp:", sp)
	fmt.Println("sp.name:", sp.name)

	sp.age = 51
	fmt.Println("sp.age:", sp.age)

	dog := struct {
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}