package main

import (
	"fmt"
	"math"
)

type person struct {
	first string
	last  string
	age   int
}

type square struct {
	lado float64
}

type circle struct {
	raio float64
}

var f func()

func main() {

	defer fooDefer()
	a, b := bar()
	fmt.Println(b, "tem", a, "anos")

	xi := []int{1, 2, 5, 6, 8, 9}
	fmt.Println("A soma total é:", foo2(xi...))
	fmt.Println("A soma total é:", foo2([]int{1, 2, 5, 6, 8, 9}...))
	fmt.Println("A soma total de xi é:", foo2(1, 2, 5, 6, 8, 9))

	xii := []int{1, 2, 3, 4, 5, 6, 7,}
	fmt.Println("A soma total de xii é:", bar2(xii))

	per := person{
		first: "Diego",
		last:  "Nogueira",
		age:   31,
	}

	per.speak()

	quadrado := square{
		lado: 3,
	}

	circulo := circle{
		raio: 5,
	}

	info(quadrado)
	info(circulo)


	func(){
		fmt.Println("Esta é uma função anonima...")
	}()

	g := func(){
		fmt.Println("Esta é uma função anonima com uma variável.")
	}
	g()
	fmt.Printf("this is g %T\n ", g)
	f = g
	f()

	re := ret()
	fmt.Println(re())
	fmt.Println(ret()())

	fmt.Println("Antigo =>", xii)
	mult := multiple(multiplicador, xii...)
	fmt.Println("Novo =>  ", mult)

	pares := foo3()
	novosPares := foo3()
	fmt.Println(pares())
	fmt.Println(pares())
	fmt.Println(pares())
	fmt.Println(pares())
	fmt.Println(novosPares())
	fmt.Println(novosPares())
	fmt.Println(novosPares())
	fmt.Println(novosPares())
	fmt.Println(foo3()())

}

func foo3() func() int {
	x := 0
	return func() int {
		x += 2
		return x
	}
}

func multiple(f func(x ...int) []int, x ...int) []int {
	return f(x...)
}

func multiplicador(x ...int) []int {
	var xi []int
	for _, value := range x {
		xi = append(xi, value*2)
	}
	return xi
}

func ret() func() int {
	return func() int {
		return 40
	}
}

func (s square) area() float64 {
	return s.lado * s.lado
}

func (c circle) area() float64 {
	return math.Pi * c.raio * c.raio
}

type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type) {
	case circle:
		fmt.Println("Área do círculo => ", s.area())
	case square:
		fmt.Println("Área do quadrado => ", s.area())
	}
}

func (p person) speak() {
	fmt.Printf("O meu nome é %s %s e tenho %d anos\n\n", p.first, p.last, p.age)
}

func fooDefer() {
	defer func() {
		fmt.Println("fooDefer defer ran")
	}()
	fmt.Println("fooDefer ran")
}

func foo() int {
	return 15
}

func bar() (int, string) {
	return 20, "Diego"
}

func foo2(x ...int) int {
	var total int
	for _, value := range x {
		total += value
	}
	return total
}

func bar2(x []int) int {
	var total int
	for _, value := range x {
		total += value
	}
	return total
}
