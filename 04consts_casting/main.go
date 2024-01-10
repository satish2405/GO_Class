package main

import "fmt"

var a int

type venkatesh int

var b venkatesh

type want_string_data string

var z venkatesh

var m want_string_data = "Hello  Sai Krishna"

func main() {
	a = 100
	b = 200
	fmt.Println(a, b, z, m)
	fmt.Printf("%T and %T", a, b)

}
