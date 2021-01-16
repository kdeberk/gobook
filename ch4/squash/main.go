package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "Hello, \t\n\u0085 W\rorld\v!"
	fmt.Println(string(squash([]byte(s))))
}

func squash(bs []byte) []byte {
	j := 0
	for i := 0; i < len(bs); {
		r, size := utf8.DecodeRune(bs[i:])
		switch {
		case unicode.IsSpace(r) && (0 == j || 0x20 != bs[j-1]):
			bs[j] = 0x20
			j++
		case !unicode.IsSpace(r):
			utf8.EncodeRune(bs[j:], r)
			j += size
		}
		i += size
	}
	return bs[:j]
}
