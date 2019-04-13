package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Aqui estáo erro: %v", ce.info)
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}






func main() {

	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("Deu merda no json", err)
	}
	fmt.Println(string(bs))

	bs, err = toJSON(p1)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))

	c1 := customErr{
		info: "Erro nessa parada",
	}

	foo(c1)


	_, err = sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}

}






func sqrt(f float64) (float64, error) {
	if f < 0 {
		//e := errors.New("Preciso de mais café")
		e := fmt.Errorf("Preciso de mais café - valor %v", f)
		return 0, sqrtError{lat:"50.2289 N", long:"99.4656 W", err:e}
	}
	return 42, nil
}

func foo(e error) {
	fmt.Println("foo ran - ", e)
	fmt.Println("foo ran - ", e.(customErr).info)
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		//return []byte{}, fmt.Errorf("Erro no json %v", err)
		return []byte{}, errors.New(fmt.Sprintf("Erro no json %v", err))
	}
	return bs, nil
}