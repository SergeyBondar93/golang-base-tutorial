package main

import (
	"fmt"
	"time"
)

/*
the second parameter of creating channel if capacity, e.g. how much values it can include before deadlock
by default 0
but if we full the channel outside from current function, and read from it, everything will be okey

if we set a capacity value to more then 1,
we can write to chanel a lot more information without lock writing
and lock will be only when capacity ends, and we will waiting for a read.
after read  readed value will be removed from the channel,
and we can write to it again and lock it again and wait for reading
*/

// func write(ch chan string) {
// 	ch <- "lol"
// 	ch <- "kek"
// }

// func main() {
// 	ch := make(chan string, 1)

// 	// go write(ch)
// 	ch <- "lol"
// 	ch <- "kek"
// 	// close(ch)

// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)

// }

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Successfully wrote ", i, " to chanel")
	}
	close(ch)
}

func main() {
	ch := make(chan int, 3)
	go write(ch)

	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("Value from channel ", v)
		time.Sleep(2 * time.Second)
	}
	fmt.Println("END")

	chnl := make(chan int, 5)
	chnl <- 5

	fmt.Println("Capacity ", cap(chnl))
	fmt.Println("Length ", len(chnl))

	chnl <- 6
	close(chnl)
	n, open := <-chnl
	fmt.Printf("Received: %v, open: %t\n", n, open)
	n, open = <-chnl
	fmt.Printf("Received: %v, open: %t\n", n, open)
	n, open = <-chnl
	fmt.Printf("Received: %v, open: %t\n", n, open)
}

// func write(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Println("successfully wrote", i, "to ch")
// 	}
// 	close(ch)
// }
// func main() {
// 	ch := make(chan int, 2)
// 	go write(ch)
// 	time.Sleep(2 * time.Second)
// 	for v := range ch {
// 		fmt.Println("read value", v, "from ch")
// 		time.Sleep(2 * time.Second)

// 	}
// }
