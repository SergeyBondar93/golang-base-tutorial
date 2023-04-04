package main

import (
	"fmt"
	"time"
)

func write(ch chan string) {
	data := []string{"My", "Name", "Is", "Sergey", "And", "I", "Am", "Golang", "Programmer"}
	for _, v := range data {
		time.Sleep(250 * time.Millisecond)
		ch <- v
		fmt.Println("Writed value: ", v)
	}
	close(ch)
}

func main() {
	ch := make(chan string, 5)

	go write(ch)

	for v := range ch {
		time.Sleep(600 * time.Millisecond)
		fmt.Println("Readed value from chanel", v)
	}

}
