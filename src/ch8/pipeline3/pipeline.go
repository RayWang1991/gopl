package main

import "fmt"

func main() {
	count := make(chan int)
	square := make(chan int)
	go counter(count)
	go squarer(count, square)
	printer(square)
}

func counter(out chan<- int) {
	for i := 0; i < 10; i ++ {
		out <- i
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
