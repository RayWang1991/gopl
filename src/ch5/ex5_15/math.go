package main

import (
	"fmt"
)

func main() {
	// let it panic
	fmt.Printf("Min: %d\nMax: %d\n", Min(1, 2, 5), Max(8, 3, 1))
}

func Min(arg int, args ...int) int {
	for _, a := range args {
		if a < arg {
			arg = a
		}
	}
	return arg
}

func Max(arg int, args ...int) int {
	for _, a := range args {
		if a > arg {
			arg = a
		}
	}
	return arg
}
