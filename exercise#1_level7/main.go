package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main() {

	x := 40
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Println("")
	fmt.Println("#######################")
	p1 := person{
		first:"Diego",
		last:"Nogueira",
		age:31,
	}
	fmt.Println("Person => ", p1)
	changeMe(&p1)
	fmt.Println("Person change => ", p1)

}

func changeMe(p *person) {
	p.first = "Lucas"
	p.last = "Telles"
	p.age = 3
	//ou isso
	//(*p).first = "Fernanda"
	//(*p).last = "Rocha"
	//(*p).age = 36
}
