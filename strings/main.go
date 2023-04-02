package main

import (
	"fmt"
	"strconv"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func main() {
	name := "Hello World"
	fmt.Println(len(name))

	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n")

	stringFromBytes()
}

func pad(n int) string {
	str := strconv.Itoa(n)

	if len(str) < 2 {
		str = "0" + str
		return str
	} else {
		return str
	}
}

func stringFromBytes() {
	var bytes []int8

	var b int8 = 0x00

	for i := 0; i < 300; i++ {
		b += 1
		bytes = append(bytes, b)
	}

	// str := string(bytes)

	fmt.Println("BYTES: ", bytes)
}
