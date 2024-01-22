package main

import (
	"fmt"
)

// var x int = 200
// var y int // zero is default value

func main() {
	var mystring = "rḛsuṁḛ"
	var indexed = mystring[1]
	fmt.Printf("%v, %T, %v, %T, length:  %d\n", mystring, mystring, indexed, indexed, len(mystring))

	for i, v := range mystring {
		fmt.Println(i, v)
	}

	for _, r := range mystring {
		runeBytes := []byte(string(r))

		fmt.Println("------------------------------")

		for _, byte := range runeBytes {
			fmt.Printf("%v \t %X \t %d \n", string(r), byte, byte)
		}
	}

	runeslice1 := []rune{0x0C05, 0x0C06, 0x002B}

	fmt.Printf("rune string slice..:%v \t %T \t %d", string(runeslice1), runeslice1, len(runeslice1))

}
