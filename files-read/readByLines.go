package main

import (
	"bufio"
	"flag"
	"fmt"
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

	s := bufio.NewScanner(f)

	for s.Scan() {
		fmt.Println(s.Text())
		time.Sleep(300 * time.Millisecond)
	}

	err = s.Err()

	if err != nil {
		log.Fatal(err)
	}

}
