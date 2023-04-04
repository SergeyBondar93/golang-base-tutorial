package main

import (
	errors "errors"
	"fmt"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("../**/*.go")
	// it will cause an error
	// files, err := filepath.Glob("../**/*.go")

	if err != nil {
		if errors.Is(err, filepath.ErrBadPattern) {
			fmt.Println("This is error named ErrBadPattern")
			return
		}
		fmt.Println("General error")
		return
	}
	fmt.Println(files)

}
