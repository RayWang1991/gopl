package count

import (
	"io"
	"fmt"
	"unicode"
	"bufio"
)

func count(reader io.Reader) (d int, l int) {
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
	return digits, letters
}
