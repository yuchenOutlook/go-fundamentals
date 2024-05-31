package my_interface_example

import "fmt"

type Programmer interface {
	Develop() string
	WriteUnitTests(testname string) string
	Debug(debugMsg string) string
}

type JavaProgrammer struct {
	Name       string
	Age        int
	Occupation string
}


type GolangProgrammer struct {
	Name string
	Address string
	Occupation string
}

func NewGolangProgrammer(name string, address string) GolangProgrammer {
	return GolangProgrammer{
		Name:       name,
		Address:    address,
		Occupation: "Golang Programmer", // Default value
	}
}

func (jp JavaProgrammer) Develop() string {
	s := fmt.Sprintf("%s is developing a new feature", jp.Name)
	fmt.Println(s)
	return s
}

func (gp GolangProgrammer) Develop() string {
	s := fmt.Sprintf("%s is developing a new feature at %s", gp.Name, gp.Address)
	fmt.Println(s)
	return s
}

func (gp GolangProgrammer) WriteUnitTests(testname string) string {
	s := fmt.Sprintf("%s is writing unit tests for %s", gp.Name, testname)
	fmt.Println(s)
	return s
}

func (gp GolangProgrammer) Debug(debugMsg string) string {
	s := fmt.Sprintf("%s is debugging %s", gp.Name, debugMsg)
	fmt.Println(s)
	return s
}