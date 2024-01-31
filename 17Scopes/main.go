package main

import (
	"fmt"

	art "github.com/satish2405/modules/articles"
)

var x int = 100 // package scope

func foo() {
	x = 400
	fmt.Println(x) // this will not work
}

func main() {

	{ // clsoures // block  scope
		i := 1
		x = 200
		fmt.Println("variable i block scope..: ", i, x)
	}
	fmt.Println("variable i outside block scope but before update value..: ", x)
	x = 300
	fmt.Println("variable i outside block scope..: ", x)
	foo()
	anotherfoo()
	//fmt.Println("Hello..! ", Name)
	fmt.Println("Here the article varaible use case..: ", art.Article_name)
}

func anotherfoo() {
	fmt.Println("another foo", x)
}
