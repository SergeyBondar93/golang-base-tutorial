package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	fptr := flag.String("filename", "file.txt", "Provide filename for reading file")

	flag.Parse()

	println(*fptr)

	p := path.Join("./", "my-files", *fptr)
	bytes, err := os.ReadFile(p)

	if err != nil {
		println("Error occurs during reading file", err)
		return
	}

	str := string(bytes)

	fmt.Println(str)

	time.Sleep(5 * time.Second)

	// er := os.WriteFile("./copy.txt", str)

}
