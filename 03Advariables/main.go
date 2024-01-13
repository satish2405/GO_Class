package main

import (
	"fmt"
)

//var y int = 10

var z string = "satish is in class"

func main() {
	var y int = 49
	var m = "sai and venky"
	r := 100
	n, _ := fmt.Println(y, z)
	fmt.Print("values of return type: ", n)
	//fmt.Printf("both values : %#v and %#v", y, z)
	//fmt.Printf("both values : %s", m)
	x, str := today_class(y, r)
	fmt.Println(x, str)

	// storing the output to a variable rather than printing on console..!!
	s := fmt.Sprintf("%v\t %#v %%", y, m)
	fmt.Println(s)
}

func today_class(y int, r int) (int, string) {
	y = 200
	fmt.Println("\nThis is upcoming session ", y)
	s := y / r
	return s, "class"
}
