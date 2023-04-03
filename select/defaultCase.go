package main

import (
	"fmt"
	"time"
)

// func process(ch chan string)  {
// 	time
// }

func process(ch chan string) {
	time.Sleep(15 * time.Second)
	ch <- "Process completed"
}

func main() {
	ch := make(chan string)

	// go process(ch)

	for {
		time.Sleep(2 * time.Second)

		select {
		case v := <-ch:
			fmt.Println("Value from channel ", v)
			return
			// with default case we prevent the deadlock statement and therefore panic
			// default:
			// 	fmt.Println("This is default")
			// 	return

		}
	}
}
