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
	specialTupleAssign()
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

func specialTupleAssign() {
	//v, ok := m[key] map lookup
	//v, ok := x.(T)  type assertion
	//v, ok := <-ch   channel recieve
	m := make(map[string]string)
	m["hello"] = "world!"
	v := m["hello"]
	fmt.Println(v)
	v, ok := m["hello"]
	fmt.Println(v)
	fmt.Println(ok)
	v, ok = m["world"]
	fmt.Println(v)
	fmt.Println(ok)
}
