package main

import "fmt"

func main() {

	//1º maneira
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println("C is", <-c)

	//2º maneira
	d := make(chan int, 1)

	d <- 35

	fmt.Println("D is", <-d)

	//2º maneira com mais de um
	e := make(chan int, 2)

	e <- 35
	e <- 45 // não funciona, só tem espaço para um

	fmt.Println("E is", <-e)
	fmt.Println("E is", <-e)
	fmt.Printf("%T", e)

	//###################################

	f := make(<-chan int, 1) //send (não posso colocar nada no canal só enviar dela para fora)

	f <- 35

	fmt.Println("F is", <-f)

	//---------------------

	g := make(chan<- int, 1) //receive (não posso retirar nada do canal só enviar para ele)

	g <- 35

	fmt.Println("G is", <-g)

}
