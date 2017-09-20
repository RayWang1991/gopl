package word2

import (
	"unicode"
)

// isPalindrome reports whether s reads the same forward and backward,
// Letter case is ignored, so as the non-letters
func IsPalindrome(s string) bool {
	var letters = make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	//len is int type
	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		if letters[i] != letters[j] {
			return false
		}
	}
	return true
}
