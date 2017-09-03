package main

import (
	"fmt"
	"bytes"
)

func main() {
	s, t := IntSet{words: []uint64{1, 4, 8}}, IntSet{words: []uint64{3, 6}}
	fmt.Printf("s: %v t:%v\n", s, t)
	fmt.Printf("s: %v t:%v\n", s.Has(2), t.Has(2))
	s.Add(2)
	fmt.Printf("s: %v t:%v\n", s.Has(2), t.Has(2))
	s.Union(&t)
	fmt.Printf("s: %s t:%s\n", &s, &t)
	t.Union(&s)
	fmt.Printf("s: %s t:%s\n", s, t)
	fmt.Print(s.String(), &t)
	fmt.Println(&s == &t)
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	if l := len(s.words); word >= l {
		s.words = append(s.words, make([]uint64, word+1-l)...)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Union(t *IntSet) {
	n, m := len(s.words), len(t.words)
	var min = m
	if n < m {
		s.words = append(s.words, t.words[n:]...)
		min = n
	}
	for i := 0; i < min; i++ {
		s.words[i] |= t.words[i]
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word != 0 {
			for j := 0; j < 64; j ++ {
				if word&(1<<uint(j)) != 0 {
					if buf.Len() > len("{") {
						buf.WriteByte(' ')
					}
					fmt.Fprintf(&buf, "%d", 64*i+j)
				}
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
