// mutex is using to avoid race conditions between goroutines

// mutex.Lock() /* ...some code */ mutex.Unlock()
// make the access to some code from only this goroutine

package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	defer m.Unlock()
	// each goroutine try to increase x cuncurrently
	// and withput lock using mutex
	// some iterations will be skipped
	// using mutex we just lock the access to variables between lock and unlock calls
	// in one goroutine we need to have only one mutex
	m.Lock()
	x++
	wg.Done()
	// we can use unlock here or with defer syntax anywhere
	// m.Unlock()
}

func main() {
	var wg sync.WaitGroup

	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &m)
	}

	wg.Wait()

	fmt.Println("X is ", x)

	fmt.Println("Everuthing is done")
}
