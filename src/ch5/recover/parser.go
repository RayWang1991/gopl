package main

import "fmt"

type AST string

func main() {
	Parse("123")
	if p := recover(); p != nil {
		fmt.Printf("Panic: %v", p)
	}
}

func Parse(input string) (s1 *AST, n int, err error) {
	//defer func() {
	//	if p := recover(); p != nil {
	//		fmt.Printf("Panic: %v", p)
	//	}
	//}()
	panic("panic 1")
	return
}
