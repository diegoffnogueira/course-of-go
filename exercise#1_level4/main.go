package main

import "fmt"

func main() {

	x := [5]int{2, 5, 7, 9, 15}

	for i, v := range x{
		fmt.Printf("indice %d com valor %d\n", i, v)
	}
	fmt.Printf("\n%T\n", x)

	y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range y{
		fmt.Printf("indice %d com valor %d\n", i, v)
	}
	fmt.Printf("\n%T\n", y)

	a := y[:5]
	fmt.Println(a)

	b := y[5:]
	fmt.Println(b)

	c := y[2:7]
	fmt.Println(c)

	d := y[1:6]
	fmt.Println(d)

	y = append(y, 52)
	fmt.Println(y)

	y = append(y, 53, 54, 55)
	fmt.Println(y)

	z := []int{56, 57, 58, 59, 60}

	y = append(y, z...)
	fmt.Println(y)

	y = append(y[:3], y[6:10]...)
	fmt.Println(y)

	w := []string{"James", "Bond", "Shaken, not stirred"}
	j := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	wj := [][]string{w, j}
	fmt.Println(wj)

	for i, vwj := range wj{
		fmt.Println("Record ", i)
		for k, val := range vwj{
			fmt.Println("Chave ", k, " Valor: ", val)
		}
	}
	
	map1 := map[string][]string{
		"bond_james": {"Shaken", "not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"`James Bond", "Literature", "Computer Science"},
		"no_dr": {"Being evil", "Ice cream", "Sunsets"},
	}

	
	
	

}
