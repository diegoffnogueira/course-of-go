package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
}

type human interface {
	speak()
}

func main() {

	p1 := person{
		first:"Diego",
		last:"Felipe",
		age:31,
	}

	sa := secretAgent{
		person: person{
			first:"Lucas",
			last:"Nogueira",
			age:2,
		},
	}

	//p1.speak()
	saySomething(&p1)
	saySomething(&sa)

}

func (p *person) speak()  {
	fmt.Printf(`%s %s says:"I have %d years"`, p.first, p.last, p.age)
}

func (s *secretAgent) speak()  {
	fmt.Printf(`O agente secreto %s %s says:"I have %d years"`, s.first, s.last, s.age)
}

func saySomething(h human) {
	h.speak()
	fmt.Println("")
	fmt.Println("-----------")
}
