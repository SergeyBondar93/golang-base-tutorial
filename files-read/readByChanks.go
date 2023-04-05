package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

func main() {
	fptr := flag.String("filename", "text.txt", "Provide filename for reading file")

	flag.Parse()

	println(*fptr)

	p := path.Join("./", "my-files", *fptr)

	f, err := os.Open(p)

	if err != nil {
		log.Fatal("Cannot open this file", err)
		return
	}
	defer func() {
		e := f.Close()
		if e != nil {
			log.Fatal("Cannot close this file", err)
		}
	}()

	r := bufio.NewReader(f)

	b := make([]byte, 3)

	for {
		_, err := r.Read(b)

		if err == io.EOF {
			fmt.Println("File fully readed")
			return
		}

		if err != nil {
			fmt.Printf("Error %s during reading file", err)
			return
		}

		fmt.Println(string(b))
		time.Sleep(300 * time.Millisecond)
	}

}
