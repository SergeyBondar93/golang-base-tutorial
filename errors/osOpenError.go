package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("../hello/ma1in.go")

	if err != nil {
		var e *os.PathError

		// errors.As resieve a pointer as a second argument
		// use it to compare types of errors and is true,
		// in next block you shpuld use a new created error (13 line)
		if errors.As(err, &e) {
			fmt.Println(e.Err)
			fmt.Println(e.Op)
			fmt.Println(e.Path)
			fmt.Println(e.Timeout())
		} else {
			fmt.Println("Unknown error")
		}
		return
	} else {
		res := f.Name()

		fmt.Println(res)

	}

}
