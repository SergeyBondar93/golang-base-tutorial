package main

import (
	"oop/composition"
	"oop/employee"
)

type website struct {
	blogs []composition.Blog
}

func main() {
	Sergey := employee.New("Bondar", "Sergey", 100000)
	Sergey.SayHi()

	author := composition.Author{FirstName: "Sergey", LastName: "Bondar"}

	post1 := composition.Blog{Title: "My live in Germany", Author: author}
	post2 := composition.Blog{Title: "My live in SPB", Author: author}
	post3 := composition.Blog{Title: "My live in USA", Author: author}
	post4 := composition.Blog{Title: "My live in Portugees", Author: author}

	siteContent := website{blogs: []composition.Blog{post1, post2, post3, post4}}

	for _, v := range siteContent.blogs {
		v.Details()
	}

}
