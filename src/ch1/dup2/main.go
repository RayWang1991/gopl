package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) > 0 {
		// read from files
		for _, args := range os.Args[1:] {
			f, error := os.Open(args)
			if error != nil {
				fmt.Fprint(os.Stderr, "dup2:%v\n", error)
			} else {
				countLine(f)
				f.Close()
			}
		}
	} else {
		// read from stdIn
		countLine(os.Stdin)
	}
}

func countLine(file *os.File) {
	count := make(map[string]int)
	input := bufio.NewScanner(file)
	for input.Scan() {
		// if count[intput.Text()] is nil, the value is treated as 0
		count[input.Text()]++
	}
	hasPrintFileName := false
	for line, n := range count {
		if n > 1 {
			if !hasPrintFileName {
				hasPrintFileName = true
				fmt.Printf("%s:\n", file.Name())
			}
			fmt.Printf("%s: %d\n", line, n)
		}
	}
}
