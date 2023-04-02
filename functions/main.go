package main

func greeting(name string) string {
	return "Hello " + name
}

func main() {
	println(greeting("Sergey"))
	println(getSum(3, 5))
}

func getSum(num1, num2 int) int {
	return num1 + num2
}
