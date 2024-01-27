package main

import "fmt"

var x int

// func increment() int {
// 	x++
// 	return x
// }

// func main() {
// 	//x := 0
// 	//var str string
// 	increment := func() int { //this is called annonymous function
// 		x++
// 		return x
// 	}

// 	fmt.Println(increment())
// 	fmt.Println(increment())
// }

// we can return a fuction as similar to like a value or variable.!
func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {

	incrementA := wrapper()
	incrementB := wrapper()

	fmt.Println("A values:", incrementA())
	fmt.Println("A values:", incrementA())
	fmt.Println("B values:", incrementB())
	fmt.Println("B values:", incrementB())
	fmt.Println("B values:", incrementB())
}
