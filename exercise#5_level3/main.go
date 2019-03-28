package main

import "fmt"

func main() {

	for i := 10; i <= 100; i++ {
		fmt.Printf("When %d is divided by 4, the remmainder is %v\n", i, i%4)
	}

}
