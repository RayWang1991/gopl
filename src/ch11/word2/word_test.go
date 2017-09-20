package word2

import (
	"testing"
	"gopl.io/ch11/word2"
)

func Test(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aaA", true},
		{"ab", false},
		// chinese test
		{"a中国中A", true},
		{"a中国中Aasdf", false},
		{"a中f", false},
		// non letter test
		{"123", true},
		{`fdf
		`, true},
	}

	for _, te := range tests {
		if got := IsPalindrome(te.input); got != te.want {
			t.Fatalf("IsPalindrome(%q)=%v, want %v", te.input, got, te.want)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func TestIsPalindrome(t *testing.T) {
	// should be
	ss := []string{"1234321", "AbaBa", "1", ""}
	for _, s := range ss {
		if !word.IsPalindrome(s) {
			t.Errorf("should be true %s \n", s)
		}
	}
}

func TestIsNotPalindrome(t *testing.T) {
	// should not be
	ss := []string{"badf", "abcaba", "asd12", "asdf"}
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
