package main

import "fmt"

func main() {
	fmt.Println("..loops..")

	for i := 0; i <= 100; i++ {
		if i%5 == 0 {
			fmt.Println(i)
		}

	}
}
