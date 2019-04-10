package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var mx sync.Mutex

	incrementor := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mx.Lock()
			v := incrementor
			v++
			incrementor = v
			fmt.Println(incrementor)
			mx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementor)

}
