package main

import (
	"sort"
	"fmt"
)

type ss string

func (s ss) Len() int {
	return len(s)
}

func (s ss) Less(i, j int) bool {
	r := []rune(string(s))
	return r[i] < r[j]
}

func (s ss) Swap(i, j int) {
	//Dont need
}

func main() {
	s := "123321"
	fmt.Println(IsPalindrome(ss(s)))
	s = "12321"
	fmt.Println(IsPalindrome(ss(s)))
	s = "121321"
	fmt.Println(IsPalindrome(ss(s)))
}

func IsPalindrome(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for i < j {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		i++
		j--
	}
	return true
}
