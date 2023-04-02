package main

import "fmt"

var name string = "Sergey"
var ageTyped int8 = 29
var age = 29 // int
var phoneTyped uint64 = 4917641234123
var phone = 491764123412311111
var isCool = true

func main() {

	// no type definition, only  in funcs and only with default type infer
	name := "Sergey"
	ageTyped := 29
	age := 29
	phoneTyped := 4917641234123
	phone := 491764123412311111
	floatVar1 := 1.5
	var floatVar2 float32 = 1.5

	isCool := true

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", ageTyped)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", phoneTyped)
	fmt.Printf("%T\n", phone)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", floatVar1)
	fmt.Printf("%T\n", floatVar2)
}
