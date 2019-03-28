package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++{
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++{
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++{
			fmt.Printf("\t\t The inner loop %d\n", j)
		}
	}

	x := 1

	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done!")


	b := 0

	for {
		b++
		if b > 100{
			break
		}
		if b % 2 != 0{
			continue
		}
		fmt.Println(b)
	}
	fmt.Println("Done!")

	for d := 33; d <= 122; d++{
		fmt.Printf("%d\t%#x\t%#U\n", d, d, d)
	}


}
