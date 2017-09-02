package main

import "fmt"

func double(x int) (result int) {
	defer func() {
		fmt.Printf("doubld(%d) = %d \n", x, result)
	}()
	return x + x
}

func triple(x int) (result int) {
	defer func() {
		result += x
	}()
	return double(x)
}

func main() {
	x := 5
	double(x)
	res := triple(x)
	fmt.Printf("The result of triple(%d) is %d", x, res)
}
