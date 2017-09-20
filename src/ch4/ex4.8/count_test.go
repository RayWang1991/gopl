package count

import (
	"testing"
	"strings"
)

func TestCount(t *testing.T) {
	tests := []struct {
		input  string
		dcount int
		lcount int
	}{
		// test number
		{"1234", 4, 0},
		{"1", 1, 0},
		{"", 0, 0},
		// test letter
		{"adfb", 0, 4},
		{"llll", 0, 4},
		{"asdff", 0, 5},
		// test non
		{"\n  &#$", 0, 0},
		{"\r\t\n", 0, 0},
		// mix
		{"\n sdfAS 中国 &#$", 0, 7},
		{"100 sdfAS 中 A", 3, 7},
	}

	for _, test := range tests {
		d, l := count(strings.NewReader(test.input))
		if d != test.dcount || l != test.lcount {
			t.Errorf("count(%q) digit%d letter%d want%d,%d", test.input, d, l, test.dcount, test.lcount)
		}
	}
}
