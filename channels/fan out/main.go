package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for value := range c2 {
		fmt.Println(value)
	}

	fmt.Println("About to exit")

}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int)  {
	var wg sync.WaitGroup
	for value := range c1 {
		wg.Add(1)
		go func(v int) {
			c2 <- timeConsumeWork(v)
			wg.Done()
		}(value)
	}
	wg.Wait()
	close(c2)
}

func timeConsumeWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(1000)))
	return n + rand.Intn(1000)
}