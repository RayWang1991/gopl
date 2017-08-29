package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	countChar(os.Stdin)
}

func countChar(reader io.Reader) {
	counts := map[rune]int{}
	lenCounts := [utf8.UTFMax+1]int{}
	invalid := 0
	scanner := bufio.NewReader(reader)
	for {
		r, l, err := scanner.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			invalid++
		} else {
			counts[r]++
			lenCounts[l]++
		}
	}
	if invalid > 0 {
		fmt.Printf("There is %d invalid rune\n", invalid)
	}
	for k, v := range counts {
		fmt.Printf("Rune:%6q %8d\n", k, v)
	}
	for i,c := range lenCounts{
		fmt.Printf("Length:%6d %8d\n", i, c)
	}
}
