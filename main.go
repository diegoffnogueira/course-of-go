package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, "says, \"Good morning James\"")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, "says, \"Shaken, not stirred.\"")
}

func saySomething(h human) {
	h.speak()
}

func main() {

	fake := struct {
		breed string
		name  string
	}{"German, Sheperd",
		"Shasta",}
	fmt.Println(fake)

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)

	m := map[string]int{
		"Diego":    31,
		"Fernanda": 36,
	}
	fmt.Println(m)

	p1 := person{
		"Diego",
		"Felipe",
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()

	saySomething(p1)
	saySomething(sa1)
}
