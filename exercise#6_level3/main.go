package main

import "fmt"

func main() {

	x := "Diego"

	if x == "Diego" {
		fmt.Println(x)
	}

	a := 4

	if a == 42{
		fmt.Println("Diferente")
	}else if a == 40 {
		fmt.Println("Igual")
	}else{
		fmt.Println("Outra coisa")
	}

	switch {
	case true:
		fmt.Println("Imprimir")
	case false:
		fmt.Println("Não imprimir")
	}

	sport := "natação"

	switch sport {
	case "natação":
		fmt.Println("Muito bom")
	case "futebol":
		fmt.Println("Ai sim foi")
	}

}
