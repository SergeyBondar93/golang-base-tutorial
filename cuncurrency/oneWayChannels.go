package main

import (
	"fmt"
	"time"
)

// pointell usage of one way channel.
// for what we need a chanel that we cannot read
// but only write?
// func sendData(sendch chan<- int) {
// 	sendch <- 10
// }

// func main() {
// 	sendch := make(chan<- int)
// 	go sendData(sendch)
// 	// we cannot read from this channel because it marked as only writable
// 	// for this we used chan<- int
// 	fmt.Println(sendch)
// }

func sendData(writeChnl chan<- int) {
	time.Sleep(1 * time.Second)
	writeChnl <- 10
	fmt.Println("This line will be executed after read from channel")
	i := 0
	for i = 0; i < 5000; i++ {
	}

	fmt.Println("Function will go untill met the next attempt to write to channel. after met it - func will wait read from channel in another place", i)

	writeChnl <- 20
	fmt.Println("This line will not be executed ever because we dont read from channel second time and value will not be recieved")
	writeChnl <- 30

	close(writeChnl)
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println("Line after gorout")
	fmt.Println("SECOND after gorout")
	// here will be pause because func will be locked until
	// some data will not be pushed to the channel we are trying to read
	v, ok := <-chnl
	fmt.Println(v, ok)
	v, ok = <-chnl
	fmt.Println(v, ok)
	v, ok = <-chnl
	fmt.Println(v, ok)
	// when ok is false that means we are trying to read from closed channel and value will be 0
	v, ok = <-chnl
	fmt.Println(v, ok)
	fmt.Println("END")
}
