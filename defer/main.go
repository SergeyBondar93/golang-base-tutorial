package main

import "fmt"

func main() {
	calc(5, 8)
	calc(9, 8)

}

func logEndOfCalculation(x int) {
	fmt.Println("End of calculation ", x)
}

func calc(a int, b int) {
	fmt.Println("START of calculation ", a, b)

	x := 0
	defer logEndOfCalculation(x)

	x = a + b
	fmt.Println("PROCESS of calculation ", a, b)
}
