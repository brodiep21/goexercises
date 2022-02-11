package interfaces

import "fmt"

type secretAgent struct {
	person
	ltk bool
}
type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first, h.(person).last)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first, h.(secretAgent).last)
	}
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

//takes in structs and interfaces and prints them for a viewing of their functionality.
func Agents() {
	b := secretAgent{
		person: person{
			"Brodie",
			"Peif",
		},
		ltk: false,
	}

	p1 := person{
		first: "Dr.",
		last:  "No",
	}

	bar(b)
	bar(p1)
}
