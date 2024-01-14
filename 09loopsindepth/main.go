package main

// 2d array, searching, sorting
// func main() {
// 	for i := 0; i < 5; i++ { //O(n^2)
// 		for j := 0; j < 5; j++ {
// 			fmt.Println(i, " - ", j)
// 		}
// 	}
// }
// 1.........50

// func main() {
// 	i := 0
// 	for {
// 		i++ // 51
// 		if i%2 == 0 { // inside
// 			continue
// 		}
// 		fmt.Printf("%d \n\n ", i) //50 doesnot print 51
//  		if i >= 50 { //51 >= 50
// 			break
// 		}
// 	}

// }

import "fmt"

func main() {

	variable := []string{"satish", "venky", "sai"} //slice

	for i, v := range variable {
		fmt.Println(i, v)
	}
}
