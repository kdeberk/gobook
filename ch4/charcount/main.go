package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF8 encodings
	invalid := 0
	var nControls, nMarks, nLetters, nPuncts, nSpaces, nSymbols int

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		switch {
		case unicode.IsControl(r):
			nControls++
		case unicode.IsMark(r):
			nMarks++
		case unicode.IsLetter(r):
			nLetters++
		case unicode.IsPunct(r):
			nPuncts++
		case unicode.IsSpace(r):
			nSpaces++
		case unicode.IsSymbol(r):
			nSymbols++
		}
	}
	fmt.Printf("rune\tcount\n")
	for r, n := range counts {
		fmt.Printf("%q\t%d\n", r, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if 0 < i {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if 0 < invalid {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	fmt.Printf("nControls:\t%d\n", nControls)
	fmt.Printf("nMarks:\t\t%d\n", nMarks)
	fmt.Printf("nLetters:\t%d\n", nLetters)
	fmt.Printf("nPuncts:\t%d\n", nPuncts)
	fmt.Printf("nSpaces:\t%d\n", nSpaces)
	fmt.Printf("nSymbols:\t%d\n", nSymbols)
}
