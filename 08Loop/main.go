package main

import "fmt"

// func main() {
// 	fmt.Println("..loops..")

// 	for i := 0; i <= 100; i++ {
// 		if i%5 == 0 {
// 			fmt.Println(i)
// 		}

// 	}
// }

// syntax same of while?
func main() {
	fmt.Println("..loops..")
	i := 1
	for i <= 100 { // remove condition for infinite
		if i%5 == 0 {
			fmt.Println(i)
			break
		}
		i++ // no increment
	}
}
