package main

import "fmt"

func main() {

	if true {
		fmt.Println("teste 001")
	}

	if false {
		fmt.Println("teste 002")
	}

	if !true {
		fmt.Println("teste 003")
	}

	if !false {
		fmt.Println("teste 004")
	}

	if 2 == 2 {
		fmt.Println("teste 005")
	}

	if 2 != 2 {
		fmt.Println("teste 006")
	}

	if x := 42; x == 142 {
		fmt.Println("100")
	}

	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print 2")
	case (3 == 3):
		fmt.Println("prints")
	case (4 == 4):
		fmt.Println("also true, does it print")
	}

	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print 2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print")
	default:
		fmt.Println("default")
	}

	switch "Bond"{
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("James Bond")
	case "Q":
		fmt.Println("Q")
	default:
		fmt.Println("default")
	}

	n := "Moneypenny"

	switch n{
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("James Bond")
	case "Q":
		fmt.Println("Q")
	default:
		fmt.Println("default")
	}

	switch n{
	case "Moneypenny", "Bodnd", "Dr No":
		fmt.Println("miss money or enyone else")
	case "Bond":
		fmt.Println("James Bond")
	case "Q":
		fmt.Println("Q")
	default:
		fmt.Println("default")
	}


}
