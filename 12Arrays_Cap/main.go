package main

import "fmt"

// len : no of elements in the slice
// max num
func main() {
	var s []int = make([]int, 3, 5) //slice another way of declaration
	//var q []int = s
	fmt.Printf("slice length and capacity: %p, %v, %v \n", s, len(s), cap(s))
	s[0] = 10
	s[1] = 20
	s[2] = 30

	fmt.Printf("slice length and capacity: %p, %v, %v \n", s, len(s), cap(s))
	//s[3] = 40
	s = append(s, 5, 6)
	fmt.Printf("slice length and capacity:%p, %v, %v \n", s, len(s), cap(s))

	s = append(s, 7, 8)
	fmt.Printf("slice length and capacity:%p, %v, %v ", s, len(s), cap(s))

}
