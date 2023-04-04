package main

import (
	"fmt"
	"runtime/debug"
)

// recover it is like error boundary or try-catch-finally
// more likely it is the catch block, cuz it has access to error message
func recoverFunc() {
	if r := recover(); r != nil {
		// in case of panic r will be a text from panic
		// after recover used the program flow will continue
		fmt.Println("recover from ", r)

		// function to don't lose the stack trace after panic
		debug.PrintStack()
	}
}

func fullName(first *string, last *string) string {
	// to stop panic in some function we should use a recover func
	// is useful only in defer mode.
	defer recoverFunc()

	// all the defers runs before panic
	// defer fmt.Println("Deffered call in fullName func")
	if first == nil {
		panic("First does not provided")
	}
	if last == nil {
		panic("Last does not provided")
	}

	return "My name is " + *first + *last

}

func sliceNormal() {
	defer recoverFunc()
	n := []int{5, 6, 7}
	fmt.Println(n[2])
	fmt.Println("Here everything is okey")

	// if we run a new goroutine, our panic recovery handlers will not work as we want
	// it works only at the same goroutines
	// not in other flows
	// go slicePanic()
	slicePanic()
}
func slicePanic() {
	n := []int{5, 6, 7}
	fmt.Println(n[5])
	fmt.Println("We access to 4 index on n o_O")
}

func main() {
	defer fmt.Println("Deffered call in main func")
	// firstName := "Sergey"
	sliceNormal()
	// fullName(&firstName, nil)

	fmt.Println("Everuthing is all right!")

}
