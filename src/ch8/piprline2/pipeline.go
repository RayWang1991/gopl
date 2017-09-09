package main

import "fmt"

func main() {
	// counter counts from 0 to 9
	num := make(chan int)
	square := make(chan int)
	go func() {
		for i := 0; i < 10; i ++ {
			num <- i
		}
		close(num)
	}()

	// squarer squares the receiving number
	go func() {
		for {
			n, ok := <-num
			if !ok {
				// the channel is closed
				close(square)
				break
			}
			square <- n * n
		}
	}()

	// pinter prints the result to std out
	for a := range square {
		fmt.Printf("%d\n", a)
	}
}
