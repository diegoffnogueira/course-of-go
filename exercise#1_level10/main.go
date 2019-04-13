package main

import "fmt"

func main() {

	//******** Exercicio 1 *************
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	a := make(chan int, 1)

	a <- 35

	fmt.Println(<-a)

	fmt.Println("")
	fmt.Println("#################################")
	//******** Exercicio 3 *************

	b := gen()
	receive(b)

	fmt.Println("about to exit")

	fmt.Println("")
	fmt.Println("#################################")
	//******** Exercicio 4 *************
	q := make(chan int)
	d := gen3(q)

	receive3(d, q)

	fmt.Println("about to exit")


	fmt.Println("")
	fmt.Println("#################################")
	//******** Exercicio 5 *************
	x := make(chan int)

	go func() {
		x <- 30
	}()

	v, ok := <- x
		fmt.Println(v, ok)

	close(x)

	v, ok = <- x
		fmt.Println(v, ok)


	fmt.Println("")
	fmt.Println("#################################")
	//******** Exercicio 6 *************

	canal := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			canal <- i
		}
		close(canal)
	}()

	for value := range canal {
		fmt.Println("Canal => ", value)
	}


	fmt.Println("")
	fmt.Println("#################################")
	//******** Exercicio 7 *************

	z := make(chan int)
	const gs = 10

	for i := 0; i < gs; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				z <- j
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println("Index:",k, "Value", <-z)
	}

}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func gen3(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive3(d, q <-chan int) {

	for {
		select {
		case v := <- d:
			fmt.Println(v)
		case <- q:
			return
		}
	}

}
