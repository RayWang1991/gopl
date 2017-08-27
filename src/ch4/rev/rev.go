package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
	testEq()
	s := make([]int, 3) // make([]int,3) returns a slice []int, len 3, cap 3
	fmt.Printf("The type is %T,len is %d, cap is %d\n", s, len(s), cap(s))
	s = s[3:]
	fmt.Printf("The type is %T,len is %d, cap is %d\n", s, len(s), cap(s))
	fmt.Println(s == nil)
	fmt.Println([0]int{} == [0]int{})
	sl := make([]int,3,5)
	sl = []int{0,1,2,3,4,5}[:3]
	fmt.Printf("The type is %T,len is %d, cap is %d\n", sl, len(sl), cap(sl))
	fmt.Print(sl[:5])
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

type A struct {
	a int
}

func testEq() {
	//a := [...]int{1, 2, 3}
	//b := [...]int{1, 2, 3}
	a, b := A{1}, A{1}
	c := [...]A{a}
	d := [...]A{b}
	fmt.Printf("is a == b ? %t\n", a == b)
	fmt.Printf("c %x %v\n", &c, c)
	fmt.Printf("d %x %v\n", &d, d)
}

// The == operation for slice is disallowed
// reason1 : the elements of slice is indirect, making it possible for a slice to contain itself (endless recursive)
// not a simple, efficient, and obvious way to deal with this
// reason2 : slice's elements is mutable, shallow == and deep == may have different behavior
