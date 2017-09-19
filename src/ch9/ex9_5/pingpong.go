package main

import (
	"time"
	"fmt"
)

func main() {
	start := time.Now()
	c := make(chan int)
	go pingpong(c)
	go pingpong(c)
	c <- 1
	time.Sleep(10 * time.Second)
	fmt.Printf("%d times in %s\n", <-c, time.Since(start))
}

// ping pong deliver the int it receives to the receiving channel
func pingpong(c chan int) {
	for {
		n := <-c
		c <- n + 1
	}
}
