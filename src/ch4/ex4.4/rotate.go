package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	rotate(s, 5)
	fmt.Println(s)
}

// rotate by n number
func rotate(s []int, n int) {
	len := len(s)
	if n < 0 || len == 0 {
		return
	}
	n = n % len
	if n == 0 {
		return
	}

	if len%n != 0 {
		// if len and n are relatively prime numbers
		// i can be [0:len]
		i := rand.Intn(len)
		t := s[i]
		for c := 0; c < len; c++ {
			j := (i + n) % len
			s[j], t = t, s[j]
			i = j
		}
	} else {
		m := len / n
		for i := 0; i < n; i ++ {
			t := s[i]
			for c := 0; c < m; c++ {
				j := (i + n) % len
				s[j], t = t, s[j]
				i = j
			}
		}
	}
}
