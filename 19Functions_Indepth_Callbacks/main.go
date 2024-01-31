package main

import "fmt"

// func callbacking(numbers ...int) {
// 	for i, v := range numbers {
// 		fmt.Println("Index and Values", i, v)
// 	}

// }

// func main() {
// 	callbacking(1, 2, 3, 4, 5)
// }

// func callbacking(numbers []int, callback func(int), callback2 func(string)) {
// 	for _, v := range numbers {
// 		//fmt.Println("Index and Values", i, v)
// 		callback(v)
// 		callback2(string(v))
// 	}

// }

// func main() {
// 	callbacking([]int{1, 2, 3, 4, 5}, func(n int) {
// 		fmt.Println("values", n)
// 	}, func(n string) {
// 		fmt.Println("stringvalues", n)
// 	})
// }

//call back with takes argument and returns arugment
// func filter(numbers []int, callback func(int) bool) []int {
// 	var xs []int //local slice

// 	for _, n := range numbers { // [1,2,3,4,5]
// 		if callback(n) { //if false ?
// 			xs = append(xs, n)
// 		}
// 	}
// 	return xs
// }

// func main() {
// 	venkatesh := filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
// 		return n > 3
// 	})

// 	fmt.Println("grater than 3 values..: ", venkatesh)
// }

// func main() {
// 	func() { //any func self exection
// 		fmt.Println("I am here..!")
// 	}()
// }

//recursive function use case.
func factorial(x int) int {

	if x == 0 {
		return 1
	}
	return x * factorial(x-1)

}

func main() {

	fmt.Println(factorial(4))
}
