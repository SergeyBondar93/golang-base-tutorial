package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("START")
	chnl := make(chan int)
	go counter(chnl)

	/* the same as below. loop with range throught channel will end when channel closes  */
	// for v := range chnl {
	// 	fmt.Println(v)
	// }
	// fmt.Println("END")

	for {
		v, ok := <-chnl
		if ok == false {
			break
		} else {
			fmt.Println(v)
		}
	}

}

func counter(chnl chan<- int) {
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		chnl <- i
	}
	close(chnl)
}
