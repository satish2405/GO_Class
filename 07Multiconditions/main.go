package main

import "fmt"

// func main() {
// 	if false {
// 		fmt.Println("first print statement")
// 	} else {
// 		fmt.Println("second print statement")
// 	}
// }

// nested if else statements
// func main() {
// 	if false {
// 		fmt.Println("first print statement")
// 	} else if false {
// 		fmt.Println("second print statement")
// 	} else if true {
// 		fmt.Println("third print statement")
// 	} else if !false {
// 		fmt.Println("fourth print statement")
// 	} else {
// 		fmt.Println("fifth print statement")
// 	}
//}

func main() {
	var logincount int = 8
	var result string

	if logincount < 10 {
		result = "User logged in"
	} else if logincount > 10 {
		result = "watched more than given times"
	} else {
		result = "Exctly watched  10 times..!"
	}
	fmt.Println(result)

	if 5/2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd number")
	}

}
