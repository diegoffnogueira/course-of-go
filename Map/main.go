package main

import "fmt"

func main() {

	x := map[string]int{
		"Diego":    31,
		"Fernanda": 36,
	}
	fmt.Println(x)
	fmt.Println(x["Diego"])
	fmt.Println(x["Lucas"])

	v, ok := x["Lucas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := x["Diego"]; ok {
		fmt.Println("This is the if print", v)
	}

	x["Lucas"] = 2

	for k, v := range x{
		fmt.Println("This is the for print", k, v)
	}

	//delete map elements
	delete(x, "Diego")
	delete(x, "Teste")
	fmt.Println(x)

	if v, ok = x["Lucas"]; ok{
		fmt.Println(v)
		delete(x, "Lucas")
	}

	fmt.Println(x)

}
