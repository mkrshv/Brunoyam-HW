package main

import (
	"fmt"
	"sync"
)

func main() {
	work := 10
	res := 0
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(work)

	go generator(ch)
	go sum(&res, &wg, ch)

	wg.Wait()

	fmt.Println(res)
}

func generator(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
}

func sum(num *int, wg *sync.WaitGroup, ch chan int) {
	for i := 0; i < 10; i++ {
		*num += <-ch
		wg.Done()
	}
}
