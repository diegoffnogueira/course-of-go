package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	first         string
	licenseToKill bool
}

func main() {

	p1 := person{
		first: "Diego",
		last:  "Nogueira",
		age:   31,
	}
	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   35,
		},
		first:         "Alguma coisa",
		licenseToKill: true,
	}

	p2 := person{
		first: "Lucas",
		last:  "Telles",
		age:   2,
	}

	fmt.Println(p1)
	fmt.Println(sa)
	fmt.Println(p2)

	fmt.Println(p1.last, p1.first, " -> ", p1.age)
	fmt.Println(sa.last, sa.person.first, sa.first, sa.licenseToKill, " -> ", p1.age)

	//anonimous struct
	an := struct {
		first string
		last  string
		age   int
	}{
		first: "Fernanda",
		last:  "Telles",
		age:   36,
	}

	fmt.Println(an)
	fmt.Println(an.age, " => ", an.first, an.last)

}
