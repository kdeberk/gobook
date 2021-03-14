package bytecounter

import (
	"bufio"
	"io"
)

type ByteCounter int64

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to bytecounter
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	start := 0
	for {
		adv, word, err := bufio.ScanWords(p[start:], true)
		start += adv
		if err != nil {
			return start, err
		}

		if word == nil {
			return start, nil
		}
		*c++
	}
}

func CounterWriter(w io.Writer) (io.Writer, *int64) {
	ctr := ByteCounter(0)
	return &ctr, (*int64)(&ctr)
}
