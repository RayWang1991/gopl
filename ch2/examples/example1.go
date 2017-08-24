package main

import (
	"flag"
	"fmt"
)

var num1 = flag.Int("n1", 0, "number arg1")
var num2 = flag.Int("n2", 0, "number arg2")
var funName = flag.String("f", "fib", "fib for fib()\ngcd for gcd()")

func main() {
	flag.Parse()
	var res int
	if *funName == "fib" {
		res = fib(*num1)
	} else {
		res = gcd(*num1, *num2)
	}
	fmt.Println(res)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
