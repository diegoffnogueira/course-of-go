package main

import "fmt"

func main() {

	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)

	y := []int{4, 5, 7, 42, 100, 102}
	fmt.Println(y)

	for i, v := range y{
		fmt.Println(i, v)
	}

	fmt.Println(y[:])
	fmt.Println(y[1:3])
	fmt.Println(y[:4])
	fmt.Println(y[2:])

	z := []int{1, 2, 3}

	y = append(y, 255, 444, 555, 666, 777)
	fmt.Println(y)

	z = append(z, y...)
	fmt.Println(z)

	//delete elements
	z = append(z[:4], z[7:]...)
	fmt.Println(z)

	//make slice
	w := make([]int, 10, 100)
	fmt.Println(w)
	w[0] = 2
	w[9] = 10
	fmt.Println(w)
	w = append(w, 15)
	fmt.Println(w)

	//slice multidimensional
	a := []string{"Diego", "Felipe", "Freitas", "Nogueira"}
	fmt.Println(a)
	b := []string{"Fernanda", "Rocha", "Telles"}
	fmt.Println(b)

	ab := [][]string{a, b}
	fmt.Println(ab)

}
