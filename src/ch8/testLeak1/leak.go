package main

import (
	"time"
	"fmt"
	"runtime"
)

// create 3 goroutines, each of them send a message to a channel after a certain time
// the main goroutine wait to receive value from that channel
// test there is any goroutine is blocked

func main() {
	test()
	panic(nil)
}

func test() {
	now := time.Now()
	done := make(chan struct{})
	for i := 1; i < 4; i++ {
		go func(i int) {
			for j := 0; j < i*1000000000; j++ {
			}
			done <- struct{}{}
		}(i)
	}
	<-done
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Printf("%.1fs pased\n", time.Since(now).Seconds())
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	fmt.Printf("%s", buf)
}
