package main

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	s := "国123中45"
	b := []byte(s)
	fmt.Printf("%s\n",b)
	reverse(b)
	fmt.Printf("%s\n",b)
}

func reverse(s []byte) {
	i, j := 0, len(s)
	for i < j {
		r1, l1 := utf8.DecodeRune(s[i:])
		r2, l2 := utf8.DecodeLastRune(s[:j])
		if r1 == utf8.RuneError || r2 == utf8.RuneError {
			fmt.Println("Decode UTF8 error!")
			return
		}
		if l1 != l2 {
			// shift
			copy(s[i+l2:j-l1], s[i+l1:j-l2])
		}
		for ri1, b1 := range []byte(string(r2)) {
			s[i+ri1] = b1
		}
		for ri2, b2 := range []byte(string(r1)) {
			s[j-l1+ri2] = b2
		}
		i += l2
		j -= l1
	}
}
