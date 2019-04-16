package main

import (
	"course-of-go/documentation/dog"
	"course-of-go/documentation/mymath"
	"fmt"
)

type cachorro struct {
	name string
	age int
}

func main() {

	fmt.Println("A soma de 2 e 3 é:", mymath.Sum(2, 3))
	fmt.Println("A soma de 10 e 20 e 30 é:", mymath.Sum(10, 20, 30))
	fmt.Println("A soma de 2 e 3 e 4 e 5 é:", mymath.Sum(2, 3, 4, 5))

	fmt.Println("")
	fmt.Println("**************************")

	lili := cachorro{
		name: "Lyli",
		age:dog.Years(10),
	}
	fmt.Println("A idade do animal é:", lili.age)

}
