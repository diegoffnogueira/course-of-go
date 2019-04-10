package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin CPUs:", runtime.NumCPU())
	fmt.Println("begin Gs:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello from thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from thing two")
		wg.Done()
	}()

	fmt.Println("mid CPUs:", runtime.NumCPU())
	fmt.Println("mid Gs:", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("exit the program")
	fmt.Println("end CPUs:", runtime.NumCPU())
	fmt.Println("end Gs:", runtime.NumGoroutine())

}
