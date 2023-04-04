package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func process(n int, wg *sync.WaitGroup) {
	fmt.Println("started goroutine ", n)

	rand := time.Duration(rand.Int31() % 10)
	fmt.Println(rand, n)

	time.Sleep(rand * time.Second)
	fmt.Println("Ended goroutine ", n)
	wg.Done()
}

func main() {
	no := 3

	var wg sync.WaitGroup

	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}

	println("This line will be printed")

	println("And this too")
	// block the main goroutine untill every in wg is done
	wg.Wait()
	println("But this will not be printed until wg group resolves")

	fmt.Println("All goroutines are ended")

}
