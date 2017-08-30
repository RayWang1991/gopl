package main

import (
	"io"
	"fmt"
	"unicode"
	"bufio"
	"os"
)

func main() {
	count(os.Stdin)
}

func count(reader io.Reader) {
	counts := map[rune]int{}
	letters, digits := 0, 0
	scanner := bufio.NewReader(reader)

	for {
		r, _, err := scanner.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error found in read from input, %v\n", err)
			return
		}
		if unicode.IsLetter(r) {
			letters++
		}
		if unicode.IsDigit(r) {
			digits++
		}
		counts[r]++
	}
	for k, v := range counts {
		fmt.Printf("rune: %6q, count%6d\n", k, v)
	}
	fmt.Println()
	fmt.Printf("letter: %6d\n", letters)
	fmt.Printf("digits: %6d\n", digits)
}
