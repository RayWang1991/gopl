package main

import (
	"bytes"
	"fmt"
)

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

func main() {
	t,s := &IntSet{}, &IntSet{}
	t.AddAll(0, 3, 7)
	s.AddAll(0,1)
	fmt.Println(s,t)
	s.IntersectWith(t)
	fmt.Println(s)

	//
	t,s = &IntSet{}, &IntSet{}
	t.AddAll(0, 3, 7)
	s.AddAll(0,1)
	s.DifferenceWith(t)
	fmt.Println(s)

	//
	t,s = &IntSet{}, &IntSet{}
	t.AddAll(0, 3, 7)
	s.AddAll(0,1)
	s.SymetricDifferenceWith(t)
	fmt.Println(s)
}

// return the number of elements
func (s *IntSet) Len() (res int) {
	for _, word := range s.words {
		for ; word != 0; word &= word - 1 {
			res ++
		}
	}
	return res
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// return a copy of te set
func (s *IntSet) Copy() *IntSet {
	w := make([]uint64, len(s.words))
	copy(w, s.words)
	return &IntSet{w}
}

// Add all from the args
func (s *IntSet) AddAll(args ... int) {
	for _, a := range args {
		s.Add(a)
	}
}

// IntersectWith
func (s *IntSet) IntersectWith(t *IntSet) {
	m := len(s.words)
	if n := len(t.words); n < m {
		m = n
	}
	for i := 0; i < m; i ++ {
		s.words[i] &= t.words[i]
	}
	s.words = s.words[:m]
}

// DifferenceWith
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, w := range s.words {
		if len(t.words) > i {
			s.words[i] = w &^ t.words[i]
		}
	}
}

// SymetricDifferenceWith
func (s *IntSet) SymetricDifferenceWith(t *IntSet) {
	for i, w := range t.words {
		if len(s.words) > i {
			s.words[i] ^= w
		} else {
			s.words = append(s.words, w)
		}
	}
}
