package main

import (
	"fmt"
)

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println("In func ", s)
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
	aa := Int{n: 44}
	fmt.Println(aa.double())
}

type Int struct {
	n int
}

func (n *Int) double() int {
	res := n.n * 2
	return res
}
