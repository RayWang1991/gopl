package main

import "fmt"

func main() {
	fmt.Println(f1())
}

func f1() (res string) {
	defer func() {
		recover()
		res = "non empty"
	}()
	panic(1)
}
