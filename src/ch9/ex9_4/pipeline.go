package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	in := make(chan int)
	wg.Add(1)
	go pipeline(in)
	in <- 1
	wg.Wait()
}

func pipeline(in chan int) {
	n := <-in
	fmt.Println(n)
	wg.Add(1)
	go pipeline(in)
	in <- n + 1
	wg.Wait()
}
