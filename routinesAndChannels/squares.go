package main

import "fmt"

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		fmt.Printf("!calcSquares %d \n", number)
		digit := number % 10
		sum += digit * digit
		number /= 10
	}

	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0

	for number != 0 {
		fmt.Printf("!calcCubes %d \n", number)
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)

	go calcCubes(number, cubech)
	go calcSquares(number, sqrch)

	squares, cubes := <-sqrch, <-cubech

	fmt.Printf("Number is %d \n", number)
	fmt.Println("Final Output", squares+cubes)

}
