package intset

import (
	"bytes"
	"fmt"
)

// wordSize is either 32 or 64 depending on the bytesize of uint.
const wordSize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represent the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/wordSize, uint(x%wordSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("IntSet{")
	for i, w := range s.words {
		buf.WriteString(fmt.Sprintf("%0*b", wordSize, w))
		if i+1 < len(s.words) {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("}")
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for 0 < word {
			word &= word - 1 // Unset last set bit
			count++
		}
	}
	return count
}

func (s *IntSet) Add(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
	for len(s.words) <= word {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		switch {
		case i < len(s.words):
			s.words[i] |= tword
		default:
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Remove(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = []uint{}
}

func (s *IntSet) Copy() *IntSet {
	ws := make([]uint, 0, len(s.words))
	copy(ws, s.words)
	return &IntSet{ws}
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *IntSet) Elems() []int {
	elems := []int{}
	for i, w := range s.words {
		n := i * wordSize
		for 0 < w {
			if 1 == 1&w {
				elems = append(elems, n)
			}
			w >>= 1
			n++
		}
	}
	return elems
}
