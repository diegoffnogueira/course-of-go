package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {

	//executa tudo de uma função e deixa esse por último
	defer boo()

	woo("Diego")
	z := bar("Vamos brincar")
	fmt.Println(z)
	x, y := zoo("Diego", "Felipe")
	fmt.Println(x)
	fmt.Println(y)

	// podemos não passar os parametros ou passar de forma diferente
	//somar()
	//somar("Diego", 2, 3, 5)
	soma := somar(2, 4, 5, 8, 9, 10, 12)
	fmt.Println(soma)
	xi := []int{1, 3, 6, 4, 8, 9, 12, 15}
	somar(xi...)
	somar()

	// podemos não passar os parametros ou passar de forma diferente
	//somar()
	//somar("Diego", 2, 3, 5)
	somarString("Dieggg")
	somarString("Dieggg", 2, 1, 3)

	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		licenseToKill: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Fernanda",
			last:  "Telles",
		},
		licenseToKill: true,
	}

	p1 := person{
		first: "Lucas",
		last:  "Telles",
	}

	//métodos em Go
	sa.speak()
	sa2.speak()
	p1.speak()
	fmt.Println("")
	
	fmt.Println("*****Interface and polimorfism******")
	poli(sa)
	poli(sa2)
	poli(p1)

	//Anonymous func
	fmt.Println("")
	fmt.Println("*****Anonymous func******")
	func() {
		fmt.Println("Estou imprimindo uma funcão sem nome")
	}()

	func(x int) {
		fmt.Println("Estou imprimindo uma funcão sem nome passando um parametro: ", x)
	}(42)

	fmt.Println("")
	fmt.Println("*****func expression******")
	f := func() {
		fmt.Println("Estou imprimindo minha primeira func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("As pessoas se encontraram no tempo de:", x)
	}
	g(1900)

	fmt.Println("")
	fmt.Println("*****return a function******")
	ret := function()
	fmt.Printf("%T\n", ret())
	//1
	retFunc := ret()
	fmt.Println(retFunc)

	//2
	fmt.Println(ret())

	//3
	fmt.Println(function()())

	fmt.Println("")
	fmt.Println("*****callback******")
	totalEven := even(somar, xi...)
	fmt.Println("total even => ", totalEven)
	fmt.Println("total odd => ", odd(somar, xi...))


	fmt.Println("")
	fmt.Println("*****closure******")
	aa := incrementor()
	bb := incrementor()
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(bb())
	fmt.Println(bb())
	fmt.Println(bb())

}

//closure
func incrementor() func() int  {
	var xx int
	return func() int {
		xx++
		return xx
	}
}
//######

//callback
func even(f func(numbers ...int) int, allInt ...int) int {
	var yi []int
	for _, value := range allInt {
		if value%2 == 0 {
			yi = append(yi, value)
		}
	}
	return f(yi...)
}

func odd(f func(numbers ...int) int, allInt ...int) int {
	var yi []int
	for _, value := range allInt {
		if value%2 != 0 {
			yi = append(yi, value)
		}
	}
	return f(yi...)
}
//#################

//return a function
func function() func() int {
	return func() int {
		return 451
	}
}

//Metodos
func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, " -> the secret agent")
}

//#######

//Interfaces
type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, " -> the person")
}

func poli(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into poliiii ", h, " and i a person")
	case secretAgent:
		fmt.Println("I was passed into poliiii ", h, " and i a secret agent")
	}
	fmt.Println("I was passed into poli ", h)
}

//#########

func boo() {
	fmt.Println("Imprimir...")
}

func woo(s string) {
	fmt.Println("Imprimiu => ", s)
}

func bar(texto string) string {
	return fmt.Sprint("Retornando => ", texto)
}

func zoo(st string, st1 string) (string, bool) {
	a := fmt.Sprint("Nome => ", st, " ", st1)
	b := false
	return a, b
}

func somar(numbers ...int) int {
	fmt.Println(numbers)
	fmt.Printf("%T\n", numbers)

	sum := 0

	for i, v := range numbers {
		sum += v
		fmt.Println("Indice ", i, " adicionando o valor ", v, " a soma ", sum)
	}
	fmt.Println("O total da soma é: ", sum)
	return sum
}

func somarString(s string, numbers ...int) int {
	fmt.Println(numbers)
	fmt.Println(s)
	fmt.Printf("%T\n", numbers)

	sum := 0

	for i, v := range numbers {
		sum += v
		fmt.Println("Indice ", i, " adicionando o valor ", v, " a soma ", sum)
	}
	fmt.Println("O total da soma é: ", sum)
	return sum
}
