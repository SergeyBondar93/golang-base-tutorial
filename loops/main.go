package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println(fizzBuzz(55)[:10])

}

func fizzBuzz(n int) []string {
	res := []string{"1"}

	for i := 2; i < n; i++ {
		by3 := i%3 == 0
		by5 := i%5 == 0

		if by3 && by5 {
			res = append(res, "FizzBuzz")
		} else if by3 {
			res = append(res, "Fizz")
		} else if by5 {
			res = append(res, "Buzz")

		} else {
			res = append(res, fmt.Sprint(i))

		}

	}

	return res
}
