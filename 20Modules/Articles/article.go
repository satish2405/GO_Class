package articles

import "fmt"

var Article_name = "venkatesh and sai go lang developer"

func Get_articles(name string) string {
	fmt.Println("printing articles in get_articles function which is articles.go file")
	return name + "articles details"
}
