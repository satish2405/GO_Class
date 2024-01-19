package main

import "fmt"

func main() {
	//keyword variable datatype
	//var venky int
	var friutsList [4]string

	friutsList[0] = "Apple"
	friutsList[1] = "pineapple"
	friutsList[2] = "banana"
	friutsList[3] = "guava"

	fmt.Println("list of all fruits", friutsList)
	fmt.Println("list of all fruits", len(friutsList), friutsList[2])

	// 1-d array
	vegetables := [...]string{"tomato", "capsicum", "potato", "onions", "chillies"}
	fmt.Println("list of vegetables in my basket: \n", vegetables, len(vegetables))

	// for _, v := range vegetables {
	// 	fmt.Println("vegetable in new line :  \n", v)
	// }
	//randomarr := [...]string{}
	for x := 0; x < len(vegetables); x++ { // 0, 1, 2, 3, 4, 5 <= 5
		fmt.Println("vegetables in new line\n", vegetables[x])
		//if x <= 3 {
		//	randomarr[x] = vegetables[x]
		//}

	}

	//2-d  array
	arr1 := [...][4]string{{"c", "python", "java"}, {"testing", "development", "production", "operations"}}
	fmt.Println("2-d array priniting..", arr1, len(arr1), len(arr1[0]))

	arr2 := [3]int{4, 5, 6}
	arr3 := [...]int{4, 5, 6}
	arr4 := [3]int{4, 5, 7}

	fmt.Println(arr2 == arr4)
	fmt.Print(arr2, arr3, arr4)

}

// sum of all array elements input: {4,5,6} == 4+5+6 = 15 === low
// second largest number in an array input: {4,5,6} = 5 is second largest element == medium
// frequency of each element in an array intput: {4, 5, 6, 5, 5} output: 4 = 1, 5 = 3, 6 = 1 == high
