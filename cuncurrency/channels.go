package main

import (
	"fmt"
	"time"
)

// data := <- a // read from the channel a and assign to variable data
// a <- data // write to the channel a from variable data
// <- a // just read the value channel

// make(chan<- int) // create a channel that we can only write in

// func main() {
// 	// channel type
// 	var a chan int
// 	if a == nil {
// 		fmt.Println("Channel a is nil , going to define it")
// 		// making a channel
// 		a = make(chan int)
// 		fmt.Printf("Type of a is %T \n", a)
// 	}
// }

func hello(done chan interface{}) {
	fmt.Println("Hello goroutine started")
	// write true to passed channel from this goroutine
	time.Sleep(3 * time.Second)

	time.Sleep(500 * time.Millisecond)
	done <- 1
	time.Sleep(500 * time.Millisecond)
	done <- 2
	time.Sleep(500 * time.Millisecond)
	done <- 3
	time.Sleep(500 * time.Millisecond)
	done <- "Almost end"
	time.Sleep(500 * time.Millisecond)

	done <- true
}

func numbers() {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}

}

func main() {
	// make a channel
	done := make(chan interface{})
	// pass this chanel as an argiment of goroutine
	fmt.Println("BEFORE HELLO")
	go hello(done)
	fmt.Println("AFTER HELLO")
	go numbers()
	fmt.Println("AFTER NUMBERS")
	var res []interface{}
	res = append(res, <-done)
	fmt.Println("ONE DONE")
	res = append(res, <-done)
	fmt.Println("TWO DONE")
	res = append(res, <-done)
	fmt.Println("THREE DONE")
	res = append(res, <-done)
	fmt.Println("LAST DONE")
	res = append(res, <-done)

	fmt.Println(done, res, "Main func ended")

}
