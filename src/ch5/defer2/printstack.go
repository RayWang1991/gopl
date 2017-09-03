package main

import (
	"runtime"
	"os"
	"fmt"
)

func main() {
	defer printStack()
	f(3)
}

func f(x int) {
	defer f(x - 1)               // can not do this cause, once x == 0 panics, defer f(0-1) is still evaluated, and cause run out of stack
	fmt.Printf("f(%d)\n", x+0/x) //panics if x == 0
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
