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
	xi := []int{1, 3, 6, 4, 8, 9}
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

	fmt.Println("")
	fmt.Println("*****Anonymous func******")
	//Anonymous func
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
