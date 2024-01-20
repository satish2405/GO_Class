package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Helloworld..!")
	var str1 string = "hello, Go"
	str2 := "\n"

	//fmt.Printf("strings. str1:  %s str2: %s and %d  \n \n", str1, str2, len(str2))
	// len ===  get the length of string in terms of BYTES not by charecters..!!
	fmt.Printf("strings. str1:  %s and %d  \n \n", str1, len(str2))

	var rawsrtring = `\n`
	fmt.Printf("strings in raw %s, %d \n", rawsrtring, len(rawsrtring))

	char := str1[2] //  it can be accessible
	fmt.Println("strings..: ", char)

	//str1[3] = char // strings are immutable so not possible to change

	str3 := "sai kumar is a great human being and helpful in nature"
	concat := str1 + "    " + str3

	fmt.Println("strings concatination ..: ", concat)

	isEqual := str1 == str3 //comparision is possible with strings

	fmt.Println("strings are equal or not:", isEqual)

	contains := strings.Contains(str3, "being")
	toUpper := strings.ToUpper(str3)
	//count := strings.Count()
	fmt.Println("is stubsting available??.", contains, toUpper)

	for _, char := range str1 { //str1 = "hello, Go"
		fmt.Printf("invidual char..:%c and %d \n\n", char, char)
	}
	// accii  1 byte
	//unicode 4 bytes
	// utf-8 code

	fmt.Println("Hello, 世界, అ 🐱‍🚀🎉🎂(●'◡'●)")
	//fmt.Printf("%+U", 😊)

	byteSlice := []byte(str1) // str1 = hello, go
	fmt.Println("byteslice..: ", byteSlice, string(byteSlice))

}
