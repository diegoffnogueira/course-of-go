package main

import "fmt"

func main() {
	fmt.Println("A soma é:", mySum(2, 10))
	fmt.Println("A soma é:", mySum(5, 6))
	fmt.Println("A soma é:", mySum(2, 3))
}

func mySum(x ... int) int {
	var sum int
	for _, value := range x {
		sum += value
	}
	return sum
}
