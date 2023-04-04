package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	// write to chanel with capacity 1 a value
	// and make some computations
	// next goroutime will not be executen until chanel is unlocked.
	// but it will be unlocked only when every current computations are done
	ch <- true
	x++
	<-ch
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, ch)
	}

	wg.Wait()

	fmt.Println("Everything is ended", x)

}
