package main

import "fmt"

const (
	a = 0 + iota
	b
	c
	d
	e
	f
	g
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
