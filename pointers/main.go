package main

import "fmt"

func main() {
	// * - get the value
	// & - get the pointer
	var a = 15

	var b = &a

	// fmt.Println(b)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T\n", a)
	// fmt.Println(*b)

	// fmt.Println(&a)

	// change value of pointer
	*b = 20

	// fmt.Println(a)

	pointer := new(int)

	fmt.Printf("Size value is %d , type is %T and address is %v \n", *pointer, pointer, pointer)

	value := 100

	pointer = &value // ( ? )the same is  *pointer = value

	fmt.Println("__AFTER CHANGE__")
	fmt.Printf("Size value is %d , type is %T and address is %v \n", *pointer, pointer, pointer)

	value = 200
	fmt.Println("__AFTER CHANGE__")
	fmt.Printf("Size value is %d , type is %T and address is %v \n", *pointer, pointer, pointer)

	changerNums()

	arrays()
	slices()
}

func changerNums() {
	num := 111

	fmt.Println(num)

	changeOnlyIntoFunction(num)

	fmt.Println(num)

	changeByPointer(&num)

	fmt.Println(num)
}

// everything in go is passed by value

func changeOnlyIntoFunction(n int) {
	n = 777

	fmt.Println("!IN FUNC!: ", n)
}
func changeByPointer(n *int) {
	*n = 9999
}

func modifyArray(arr *[3]int) {
	arr[0] = 999
}

// not a good way to modify an array
func arrays() {
	array := [3]int{4, 5, 6}

	modifyArray(&array)

	fmt.Println(array)

}
func modifySlice(s []int) {
	s[0] = 777
}

func slices() {
	slice := []int{7, 8, 9}

	modifySlice(slice)

	/*
		we can do the same if we write
		modifySlice(slice...)
		and in modifySlice function
		func modifySlice(s ...int)
	*/

	fmt.Println(slice)

}
