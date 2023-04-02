package main

import "fmt"

func main() {
	// Arrays
	var names [2]string

	names[0] = "Sergey"
	names[1] = "Anna"

	fmt.Println(names, names[0], names[1])

	fruitsArray := [2]string{"Apple", "Orange"}

	fmt.Println(fruitsArray)
	fmt.Println(len(fruitsArray))
	fmt.Printf("%T\n", fruitsArray)

	fruitsSlice := []string{"Apple", "Orange", "Banana"}

	fmt.Println(fruitsSlice)
	fmt.Println(fruitsSlice[:2])
	fmt.Printf("%T\n", fruitsSlice)
}
