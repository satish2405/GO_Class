package main

import "fmt"

const (
	c0 = iota // 0
	c1 = iota // 1
	c2 = iota // 2
)

const (
	c3 = iota
	c4
	c5
	c6
)

func main() {
	//c = 20
	fmt.Println(c0, c1, c2)
	fmt.Println(c3, c4, c5, c6)
}
