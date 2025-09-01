package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go add(ch, &wg)
	go get(ch)
	wg.Wait()
}

func add(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func get(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("received: %d\n", <-ch)
	}
}
