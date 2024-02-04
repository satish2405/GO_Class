package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
)

// user-defined datatype contains collection of named fileds
type Adress struct {
	Name   string
	Street string
	City   string
	State  string
	Age    int32 //newly added..!
}

type Person struct {
	FirstName string
	Adress    //embedded structs
}

func main() {
	result1 := rand.Int63()
	result2 := rand.Int63()

	complex_number := cmplx.Abs(3 + 4i)

	fmt.Println("Random Numbers :", result1)
	fmt.Println("Random Numbers :", result2)
	fmt.Println(complex_number)

	p := Person{
		FirstName: "Satish",
		Adress: Adress{
			Name:   "Kumar",
			Street: "Begumpet",
			City:   "Hyderabad",
			State:  "Telangana",
		},
	}

	fmt.Println("First Name..:", p.FirstName)
	fmt.Println("First Name..:", p.Adress)
	fmt.Println("First Name..:", p.State)
}
