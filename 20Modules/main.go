package main

import (
	"fmt"

	art "company.com/events/Articles"
	sub "company.com/events/Articles/subarticals"
)

func main() {
	fmt.Println("Hello..! ", name)
	get_Role("Lead Developer")
	fmt.Println("Call Articles as module.:", art.Article_name)
	art.Get_articles("Lifeissail")
	fmt.Println("import nested package  ..: ", sub.WelcomeText)
}
