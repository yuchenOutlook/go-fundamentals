package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName   xml.Name	`xml:"plant"`
	Id			int		`xml:"id,attr"`	
	Name		string	`xml:"name"`
	Origin		[]string	`xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	p := fmt.Println
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// out is slice of bytes
	out, _ := xml.MarshalIndent(coffee, " ", "  ")

	p("---------------")
	p("XML printed without header: ")
	p(string(out))
	p("===========================")
	p("XML printed WITH header:")
	p(xml.Header + string(out))
	p()
	
	p("Print the unmarhsed results: ")
	p("==============================")
	var pl Plant
	if err := xml.Unmarshal(out, &pl); err != nil {
		panic(err)
	}
	p(pl)
	p("===========================")
	p()

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName		xml.Name	`xml:"nesting"`
		Plants		[]*Plant	`xml:"parentItem>childItem>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	p()
	p("Try to print the nesting object")
	p("\n============================")
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	p(string(out))

	
}