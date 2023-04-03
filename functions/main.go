package main

func greeting(name string) string {
	return "Hello " + name
}

type myInt int

func (n myInt) add(n2 myInt) myInt {
	sum := n + n2

	return sum
}

func main() {
	// println(greeting("Sergey"))
	// println(getSum(3, 5))

	// next two lines are the same
	// var a myInt = 15
	a := myInt(15)

	b := a.add(30)

	println(b)

}

func getSum(num1, num2 int) int {
	return num1 + num2
}
