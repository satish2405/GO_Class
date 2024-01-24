package main

import (
	"fmt"
)

// declaring a struct
type Person struct {
	Firstname string
	Lastname  string //""
	Age       int
}

func main() {
	//Declare a  variable
	var p Person
	fmt.Println("person details..:", p) // default values

	p1 := Person{"Sai", "Venkat", 25}
	fmt.Println("Person..P1", p1)

	p2 := Person{
		Firstname: "Johney",
		Lastname:  "Canady",
		Age:       45,
	}
	fmt.Println("Person. P2", p2)

	p3 := Person{Firstname: "RAMA"}
	fmt.Println("Person P3:", p3)

	//accessing struct field
	fmt.Println("Person1 Firsname: ", p1.Firstname)
	fmt.Println("Person2 Age: ", p2.Age)

	p1.Firstname = "Bill Gates" // member only updated
	fmt.Println("P1 updated details..:", p1)

	p2 = Person{Firstname: "Ramakrishna"} // full updated
	fmt.Println("Person2 updated as a whole.:", p2)

	newperson := Person{"Jack", "John", 12}

	// pointer to the person struct
	pp := &newperson
	fmt.Println(pp)

	//accessing struct fiels using pointer
	fmt.Println("Values find using pointer..:", (*pp).Firstname)
	//fmt.Println("Values find using pointer..:", pp.Firstname)

	pp.Age = 40
	fmt.Println("pp is updated?", pp, newperson)

	//new way of declation to get pointer to struct using build-in new() function
	nPerson := new(Person)

	nPerson.Age = 100
	nPerson.Firstname = "Sachin"

	fmt.Println(nPerson)

	p5 := Person{"Jack", "Rakesh", 40}
	fmt.Println(p5)

	if p5 == newperson {
		fmt.Println("Point p5 and newperson is equal")
	} else {
		fmt.Println("Point p5 and newperson is not equal")
	}

}
