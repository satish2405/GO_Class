//inheritece by using methods.

//func (variable_name Datatype) functionname(arguments) return_type == method declartion

// func (venky int) sum(a , b int) int { } == method declartion example.

// package main

// import "fmt"

// type Point struct {
// 	x, y float64
// }

// func Isgrater(p Point, g float64) bool { //fuction definition with return bool
// 	return p.y > g // 3.3 > 2
// }

// func main() {
// 	p := Point{2.2, 3.3}

// 	fmt.Println("Point Values inside ..: ", p)
// 	fmt.Println("method call output : ", Isgrater(p, 2)) //fuction  call with arguments
// }

// package main

// import "fmt"

// type Point struct {
// 	x, y float64
// }

// func (p Point) Isgrater(g float64) bool { //method definition with return bool
// 	return p.y > g // 3.3 > 2
// }

// func main() {
// 	p := Point{2.2, 3.3}

// 	fmt.Println("Point Values inside ..: ", p)
// 	fmt.Println("method call output : ", p.Isgrater(2)) //method  call with arguments
// }

package main

import (
	"fmt"
	"math"
)

type ArthemicProgression struct {
	A, D float64
}

type GeometricProgression struct {
	A, R float64
}

func (ap ArthemicProgression) findaterm(n int) float64 {
	return ap.A + ap.D
}

func (gp GeometricProgression) findaterm(n int) float64 {
	return gp.A * math.Pow(gp.R, 2)
}

func main() {
	ap := ArthemicProgression{1, 2}
	gp := GeometricProgression{1, 3}

	fmt.Println("arthemetic one is : ", ap.findaterm(4))
	fmt.Println("arthemetic one is : ", gp.findaterm(4))
}
