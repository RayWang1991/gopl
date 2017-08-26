package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "12345678"
	c0 := comma0(s)
	c1 := comma1(s)
	fmt.Println(c0)
	fmt.Println(c1)
	s = "12345678.1234"
	c2 := comma2(s)
	fmt.Println(c2)
}

// comma inserts commas in a non-negative decimal integer string
func comma0(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma0(s[0:n-3]) + "," + s[n-3:]
}

// comma non-recursive version
func comma1(s string) string {
	return commaBackword(s, 3)
}

// comma non-recursive version, supporting number (int or float) with optional sign
func comma2(s string) string {
	cI := strings.Index(s, ".")
	if cI < 0 {
		return commaBackword(s, 3)
	}
	sf := s[0:cI]
	sb := s[cI+1:]
	return commaBackword(sf, 3) + "." + commaForward(sb, 3)
}

func commaBackword(s string, sepn int) string {
	n := len(s)
	for n > sepn {
		s = s[0:n-sepn] + "," + s[n-sepn:]
		n -= sepn
	}
	return s
}

func commaForward(s string, sepn int) string {
	len := len(s)
	for len > sepn {
		s = s[0:sepn] + "," + s[sepn:]
		len++
		sepn += sepn + 1
	}
	return s
}
