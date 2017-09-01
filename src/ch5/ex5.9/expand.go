package main

import (
	"strings"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: enter a string")
	}
	s := os.Args[1]
	fmt.Println(expand(s,
		func(foo string) string {
			return foo + foo
		}))
}

func expand(s string, f func(string) string) string {
	// found sub string $foo
	i := strings.Index(s, "$")
	if i < 0 {
		return s
	}
	l := len("$")
	var sub string
	// result
	res := string(s[:i])
	// remain undispoed string
	// can use split instead
	s = string(s[i+l:])
	for i >= 0 {
		// find sub
		i = strings.Index(s, "$")
		if i < 0 {
			sub = s
		} else {
			sub = string(s[:i])
			s = string(s[i+l:])
		}
		if f != nil {
			sub = f(sub)
		}
		res += sub
	}
	return res
}
