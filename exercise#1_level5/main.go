package main

import (
	"fmt"
)

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	p1 := person{
		firstName:       "Diego",
		lastName:        "Felipe",
		iceCreamFlavors: []string{"chocolate", "morango", "flocos", "napolitano"},
	}

	p2 := person{
		firstName:       "Fernanda",
		lastName:        "Telles",
		iceCreamFlavors: []string{"Creme", "Baunilha", "Laranja", "Uva"},
	}

	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.iceCreamFlavors {
		fmt.Println("	", v)
	}

	fmt.Println("")
	fmt.Println("######")
	fmt.Println("")

	fmt.Println(p2.firstName, p2.lastName)
	for _, v2 := range p2.iceCreamFlavors {
		fmt.Println("	", v2)
	}

	fmt.Println("")
	fmt.Println("######")
	fmt.Println("")

	mapP1 := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(mapP1)

	for _, val := range mapP1 {
		fmt.Println(val.firstName, val.lastName)
		for i, val2 := range val.iceCreamFlavors {
			fmt.Println(i, "	", val2)
		}
	}

	fmt.Println("")
	fmt.Println("######")
	fmt.Println("")

	tr := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Blue",
		},
		fourWheel: false,
	}

	sd := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "Red",
		},
		luxury: true,
	}

	fmt.Println(tr)
	fmt.Println(sd)
	fmt.Println("Truck ", tr.doors, " color ", tr.color, " => ", tr.fourWheel)
	fmt.Println("Sedan ", sd.doors, " color ", sd.color, " => ", sd.luxury)

	fmt.Println("")
	fmt.Println("######")
	fmt.Println("")

	st := struct {
		nome    string
		amigos  map[string]int
		bairros []string
	}{
		nome: "Diego",
		amigos: map[string]int{
			"Carlos":   36,
			"Eduardo":  17,
			"Anderson": 32,
		},
		bairros: []string{"Bangu", "Madureira", "Valqueire", "Realengo"},
	}

	fmt.Println(st)

	for chave, valor := range st.amigos {
		fmt.Println(chave, " => ", valor)
	}

	for chave, valor := range st.bairros {
		fmt.Println(chave, " => ", valor)
	}

}
