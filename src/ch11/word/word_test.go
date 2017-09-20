package word

import (
	"testing"
	"gopl.io/ch11/word1"
)

func TestIsPalindrome(t *testing.T) {
	// should be
	ss := []string{"1234321", "ababa", "1", ""}
	for _, s := range ss {
		if !word.IsPalindrome(s) {
			t.Errorf("should be true %s \n", s)
		}
	}
}

func TestIsNotPalindrome(t *testing.T) {
	// should not be
	ss := []string{"12134321", "abcaba", "12", "asdf"}
	for _, s := range ss {
		if word.IsPalindrome(s) {
			t.Errorf("should [not] be true %s \n", s)
		}
	}
}

// UTF8 encoded strings should be satisfied
func TestIsChinesePalindrome(t *testing.T) {
	// should not be
	ss := []string{"中", "ab中哈中ba",}
	for _, s := range ss {
		if !word.IsPalindrome(s) {
			t.Errorf("should be true %s \n", s)
		}
	}
}
