package main

import (
	"fmt"
	"os"
)

var text1 = "Hello world!"
var text2 = string([]byte{97, 97, 97, 97, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100})

func main() {
	f, err := os.Create("one.txt")

	if err != nil {
		fmt.Println("Error create file", err)
		return
	}

	l, err := f.WriteString(text2)

	if err != nil {
		fmt.Println("Error write file", err)
		f.Close()
		return
	}

	fmt.Println(l, " Bytes writed successfully")

	err = f.Close()

	if err != nil {
		fmt.Println(err)
	}

}
