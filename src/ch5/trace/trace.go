package main

import (
	"time"
	"fmt"
)

func main() {
	defer trace("operation")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	now := time.Now()
	return func() {
		fmt.Printf("%s start: %v end:%v\n", msg, now, time.Now())
	}
}
