package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"sort"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type SecretAgent struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

type ByAge []Person

type ByName []Person

func main() {

	p1 := Person{
		First: "Diego",
		Last:  "Felipe",
		Age:   31,
	}

	p2 := Person{
		First: "Lucas",
		Last:  "Telles",
		Age:   2,
	}

	p3 := Person{
		First: "Carlos",
		Last:  "Nobre",
		Age:   40,
	}

	p4 := Person{
		First: "Fernanda",
		Last:  "Souza",
		Age:   36,
	}

	people := []Person{
		p1,
		p2,
		p3,
		p4,
	}
	fmt.Println("")
	fmt.Println("########Marshal########")

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error =>", err)
	}
	fmt.Println(string(bs))

	fmt.Println("")
	fmt.Println("########UnMarshal########")

	str := `[{"First":"Diego","Last":"Felipe","Age":31},{"First":"Lucas","Last":"Telles","Age":2}]`
	bts := []byte(str)
	var secretAgent []SecretAgent
	err = json.Unmarshal(bts, &secretAgent)
	if err != nil {
		fmt.Println("Error =>", err)
	}
	fmt.Println(secretAgent)
	for key, value := range secretAgent {
		fmt.Println("Index key =>", key)
		fmt.Println("Value =>", value.First, value.Last, value.Age)
	}

	fmt.Println("")
	fmt.Println("########SORT SLICE########")
	xi := []int{2, 8, 6, 7, 12, 45, 78, 9,}
	xs := []string{"Diego", "Aline", "Carlos", "Fernanda", "Lucas"}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println("------------------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Println("")
	fmt.Println("########SORT SLICE Struct########")
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	fmt.Println("------------------")
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

	fmt.Println("")
	fmt.Println("########bcrypt########")
	s := "password123"
	bs, err = bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error =>", err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs))
	fmt.Println("--------------")
	loginPassword := "password123"
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println("Você não pode logar!!")
	} else {
		fmt.Println("Parabéns... Você está logado!!")
	}

}

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (bn ByName) Len() int {
	return len(bn)
}
func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}
func (bn ByName) Less(i, j int) bool {
	return bn[i].First < bn[j].First
}
