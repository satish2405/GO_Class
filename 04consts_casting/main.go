package main

import "fmt"

var a int

type venkatesh int

var b venkatesh

type want_string_data string
type want_int_data int64

var z venkatesh

var m want_string_data = "Hello  Sai Krishna"

var n want_int_data = 1234

func main() {
	a = 100
	b = 200
	fmt.Println(a, b, z, m, n)
	fmt.Printf("%T and %T", a, b)

}
