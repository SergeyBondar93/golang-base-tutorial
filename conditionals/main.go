package main

import "fmt"

func main() {
	x := 100
	y := 150

	color := "red"

	if x < y {
		fmt.Println("X Less then Y")
	} else if x == y {
		fmt.Println("Y equals to X")
	} else {
		fmt.Println("Y Less then X")

	}

	switch color {
	case "blue":
		fmt.Println("Color is blue")
	case "green":
		fmt.Println("Color is green")
	case "red":
		fmt.Println("Color is red")
	default:
		fmt.Println("Unrecognized color")
	}
}
