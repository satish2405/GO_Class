package main

import "fmt"

func printArray(arr [10]int) int {
	var result int
	for index, value := range arr {
		fmt.Println(index, value)
		result += value
	}
	return result
}

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //lighiting operator

	b := a //  copy of the array by value
	a[2] = 300
	fmt.Println("array a output ..:", a)
	fmt.Println("array b output ..:", b)

	refb := &a // copy of array by reference
	//a[8] = 900
	refb[2] = 200
	fmt.Println("array a output ..:", a)

	x := a[:len(a)-1]
	fmt.Println(x)

	y := a[5:]
	fmt.Println(y)

	z := a[:] //total array copy
	fmt.Print(z)

	m := a[2:6]
	fmt.Println(m)

	newvar := printArray(a)

	fmt.Println(newvar)

	//var s []int = []int{1, 2, 3} //slice

	var s []int = make([]int, 3, 5) //slice another way of declaration
	var q []int = s
	q[0] = 10
	fmt.Println("Values of S: ", s)
	fmt.Println("Values of Q: ", q)

	fmt.Println(len(s))
	fmt.Println(cap(q))

	var g = []int{1, 2, 3, 4}
	var h = append(g, 5, 6, 7)
	fmt.Println(g, h)

	//var i = append(h, g[0], g[1], g[2], g[3])
	var i = append(h, g...)
	fmt.Println("values in i : ", i)
}
