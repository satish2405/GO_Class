package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type User struct {
// 	Id         int    `json: "id"`
// 	Name       string `json: "name"`
// 	Occupation string `json: "occupation"`
// }

// func main() {
// 	user := User{Id: 1, Name: "Satish Kumar", Occupation: "Software Engineer"}
// 	result, _ := json.MarshalIndent(user, " ", " ") //tap space (to print neatly)
// 	fmt.Println(string(result))
// }

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	Id         int    `xml: "id"`
	Name       string `xml: "name"`
	Occupation string `xml: "occupation"`
}

func main() {
	user := User{Id: 1, Name: "Satish Kumar", Occupation: "Software Engineer"}
	result, readingerror := xml.MarshalIndent(user, " ", " ") //tap space (to print neatly)
	fmt.Println(result, readingerror)
	fmt.Println(string(result))
}

// task to do decoding hint: Unmarshal
