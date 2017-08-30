package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
)

func main() {
	countWord(os.Stdin)
}

func countWord(reader io.Reader) {
	counts := map[string]int{}
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		counts[text]++
	}
	for k, v := range counts {
		fmt.Printf("word: %q, fre: %d\n", k, v)
	}
}
