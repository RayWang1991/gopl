package main

import "fmt"

type st struct {
	a, b int
	c    string
}

func main() {
	testStruct()
}

func testStruct() {
	s1, s2 := st{1, 1, "c"}, st{1, 1, "c"}
	fmt.Println(s1 == s2)
	ps1, ps2 := &st{1, 1, "c"}, &st{1, 1, "c"}
	fmt.Println(ps1 == ps2)
}
