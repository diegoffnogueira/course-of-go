package main

import (
	"fmt"
	"sync"
)

func main() {

	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanIn)

	for value := range fanIn {
		fmt.Println(value)
	}

	fmt.Println("About to exit")

}

func send(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0{
			e <- i
		}else{
			o <- i
		}
	}
	close(e)
	close(o)
}

func receive(e, o <-chan int, fanin chan<- int)  {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for value := range e {
			fanin <- value
		}
		wg.Done()
	}()

	go func() {
		for value := range o {
			fanin <- value
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
