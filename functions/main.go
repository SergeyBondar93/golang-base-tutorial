package main

import (
	"fmt"
)

func greeting(name string) string {
	return "Hello " + name
}

type myInt int

func (n myInt) add(n2 myInt) myInt {
	sum := n + n2
	return sum
}

type sum func(a int, b int) int

var s sum = func(a int, b int) int {
	return 123
}

func main() {
	// println(greeting("Sergey"))
	// println(getSum(3, 5))

	// next two lines are the same
	// var a myInt = 15
	a := myInt(15)

	b := a.add(30)

	println(b)

	fmt.Println(fibo(10))
}
func getSum(num1, num2 int) int {
	return num1 + num2
}

func fibo(n int) int {
	res := []int{0, 1, 1}

	for i := 2; i < n; i++ {
		res = append(res, res[len(res)-1]+res[len(res)-2])
	}
	return res[len(res)-1]
}
