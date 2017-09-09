package main

import (
	"time"
	"fmt"
)

func main() {
	go spinner(100 * time.Millisecond)
	n := 45
	fibN := fib(n)
	fmt.Printf("\r Fib(%d) is %d\n", n, fibN)
}

func spinner(d time.Duration) {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(d)
		}
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
