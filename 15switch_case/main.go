package main

import "fmt"

// multiple alternatives we have to  choose  decision based on recived info
// func main() {
// 	switch "satish" {
// 	case "kumar":
// 		fmt.Println("Hai Kumar")
// 	case "krishna", "sai":
// 		fmt.Println("Hai sai")
// 		fallthrough //we should always keep at last
// 	case "srinu":
// 		fmt.Println("Hai srinu") // we can add mutiple logic here
// 		fmt.Println("He is doing well..!")
// 	case "satish":
// 		fmt.Println("Hai satish")
// 		fallthrough
// 	default:
// 		fmt.Println("None of them found")
// 	}
// }

// func main() {
// 	myfriendname := "venkat"

// 	switch {
// 	case len(myfriendname) == 2:
// 		fmt.Println("Yes his name length is just 2 char")
// 	case myfriendname == "kumar":
// 		fmt.Println("Hai Kumar")
// 	case myfriendname == "krishna":
// 		fmt.Println("Hai sai")
// 		fallthrough //we should always keep at last
// 	case myfriendname == "srinu":
// 		fmt.Println("Hai srinu") // we can add mutiple logic here
// 		fmt.Println("He is doing well..!")
// 	case myfriendname == "satish":
// 		fmt.Println("Hai satish")
// 		fallthrough
// 	default:
// 		fmt.Println("None of them found")
// 	}
// }

//template, class
type contact struct {
	greeting string
	id       int
	name     string
}

func swithontype(x interface{}) {
	switch x.(type) { // this is assert, asserting
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float")
		//fallthrough
	case contact:
		fmt.Println("Our Own structure contact")
		//fallthrough
	default:
		fmt.Println("Unknown Type")
	}

}
func main() {

	//swithontype(10)
	//swithontype("HaiVenkat")
	//swithontype(2.2)
	var t = contact{"Good to see you", 1234, "Kumar"}
	swithontype(t)
	swithontype(t.greeting)
	//fmt.Println("structure member...:", t.greeting, t.id, t.name)
	//swithontype(t.id)

}
