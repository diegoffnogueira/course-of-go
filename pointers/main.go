package main

import "fmt"

func main() {

	a := 42
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // ma passa o endereço onde está a variável

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // me dá o valor da variável armazenada no endereço
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)

	//use pointers
	fmt.Println("")
	fmt.Println("######No pointer########")
	x := 0
	foo(x)
	fmt.Println(x)

	fmt.Println("")
	fmt.Println("######use pointer########")
	z := 0
	fmt.Println(&z)
	fmt.Println(z)
	fooPointer(&z)
	fmt.Println(&z)
	fmt.Println(z)
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func fooPointer(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 35
	fmt.Println(y)
	fmt.Println(*y)
}
