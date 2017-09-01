package main

import "fmt"

// anonymous func, closure
// the anonymous func retain the local var n by addr?
func squares() func() int {
	var n int
	return func() int {
		n++
		return n * n
	}
}

func main () {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
