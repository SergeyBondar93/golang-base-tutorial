package main

import (
	_ "embed"
	"fmt"
)

//go:embed my-files/text.txt
var contents []byte

func main() {
	fmt.Println("Contents of file:", string(contents))
}
