package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4}
	fmt.Println(a)
	reverse(&a)
	fmt.Println(a)
}

func reverse(intArr *[5]int) {
	for i, j := 0, len(intArr)-1; i < j; i, j = i+1, j-1 {
		intArr[i], intArr[j] = intArr[j], intArr[i]
	}
}
