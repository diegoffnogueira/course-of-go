package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type User struct {
	First string `json:"Nome"`
	Last  string `json:"Sobrenome"`
	Age   int    `json:"Idade"`
}

type Person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type ByAge []Person

type ByName []Person

func main() {

	u := User{
		First: "Diego",
		Last:  "Felipe",
		Age:   31,
	}

	u2 := User{
		First: "Lucas",
		Last:  "Telles",
		Age:   2,
	}

	u3 := User{
		First: "Fernanda",
		Last:  "Rocha",
		Age:   36,
	}

	users := []User{
		u,
		u2,
		u3,
	}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Erro!")
	}
	fmt.Println(string(bs))

	fmt.Println("--------------------------------")

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var person []Person

	err = json.Unmarshal([]byte(s), &person)
	if err != nil {
		fmt.Println("Erro!", err)
	}

	for _, value := range person {
		fmt.Printf("%s%s tem %d anos e says:\n", value.First, value.Last, value.Age)
		for _, v := range value.Sayings {
			fmt.Printf("\t\t%s\n", v)
		}
		fmt.Println("")
	}

	fmt.Println("------------------")
	err = json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}

	fmt.Println("------------------")


	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println("Sem ordem",xi)
	sort.Ints(xi)
	fmt.Println("Com ordem", xi)
	fmt.Println("")
	fmt.Println("Sem ordem", xs)
	sort.Strings(xs)
	fmt.Println("Com ordem", xs)


	fmt.Println("------------------")
	fmt.Println("Original =>", person)

	sort.Sort(ByAge(person))
	fmt.Println("By age => ", person)
	for _, value := range person {
		fmt.Println(value.First, value.Last, value.Age)
			sort.Strings(value.Sayings)
		for _, val := range value.Sayings {
			fmt.Println("		", val)
		}
	}

	sort.Sort(ByName(person))
	fmt.Println("By name => ", person)

	
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
	return bn[i].Last < bn[j].Last
}
