package main

import (
	"fmt"
	"sync"
)

func main() {
	work := 5
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(work)
	go sender(ch)
	go reciever(&wg, ch)
	wg.Wait()
}

func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func reciever(wg *sync.WaitGroup, ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Print(<-ch, " ")
		wg.Done()
	}

}
