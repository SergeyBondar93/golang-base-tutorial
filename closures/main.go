package main

import "fmt"

func main() {
	adder := sum(15)

	sum := adder(10)

	fmt.Println(sum)

	sum = adder(15)

	fmt.Println(sum)

}

func sum(init int) func(int) int {
	var sum int = init

	return func(b int) int {
		sum += b
		return sum
	}
}
