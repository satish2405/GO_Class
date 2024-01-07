package main

import "fmt"

func main() {
	var a int                                                      //long declaration
	var b float32 = 10.009                                         //long declaration
	var c string = "venkatesh"                                     //long declaration
	c = "satish"                                                   //short declaration
	d := "sai"                                                     //short declaration
	e := 20.90909                                                  //short declaration
	fmt.Println("Hello Variables..!", a, b, c, d)                  //values printing
	fmt.Printf("Name of variable : %s and datatype : %T \n", d, d) //printing the data type as well
	fmt.Printf("Name of variable : %f and datatype : %T \n", e, e)

	var x, y, z int = 2, 5, 10 // multi long decalartion in single line
	var (                      // multidata long type declaration
		name   string = "sai kumar"
		age    int
		weight float64 = 67.89
	)
	fmt.Println("Hello Variables..!", x, y, z)
	fmt.Println(name, age, weight)

	m, n, o := 2, 5.2, "satish"
	fmt.Println(m, n, o)
	fmt.Printf("datatypes : %T, %T, %T \n", m, n, o)

}

// long declaration, short delcaration, muliple short and long declaration

// int, float, boolean, char, switch
// struct, union, enum, typedef
