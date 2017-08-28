package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("1中2中")
	fmt.Println(string(b))
	b = squash(b)
	fmt.Println(string(b))
}

func squash(bs []byte) []byte {
	out := bs
	j := 0
	for i := 0; i < len(bs); {
		r, l := utf8.DecodeRune(bs[i:])
		if r == utf8.RuneError {
			//
			fmt.Printf("utf8 decode error: %v", bs[i:])
			return out
		}
		//if unicode.IsSpace(r) {
		if r == '中' {
			out[j] = ' '
			out = out[:len(out)+1-l ]
			j ++
		} else {
			copy(out[j:j+l], bs[i:i+l])
			j += l
		}
		i += l
	}
	return out
}
