package composition

import "fmt"

type Blog struct {
	Title, Content string
	Author         Author
}

func (b Blog) Details() {
	fmt.Println(
		"Title: ",
		b.Title,
		" Content ",
		b.Content,
		" Author ",
		b.Author.sayHi())
}
