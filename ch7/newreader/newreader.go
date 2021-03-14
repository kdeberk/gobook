package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	r := NewReader("<html><h1>Foo</h1><h2>Bar</h2></html>")
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elementcount: %v", err)
		os.Exit(1)
	}

	counts := make(map[string]int)
	visit(counts, doc)
	for el, count := range counts {
		fmt.Println(el, count)
	}
}

func visit(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(counts, c)
	}
}

type reader struct {
	start int
	s     []byte
}

func (r *reader) Read(p []byte) (n int, err error) {
	switch {
	case len(r.s) <= r.start:
		err = io.EOF
	case len(r.s) <= r.start+len(p):
		n = len(r.s) - r.start
		copy(p, r.s[r.start:])
		r.start += n
	default:
		n = len(p)
		copy(p, r.s[r.start:])
		r.start += n
	}
	return
}

func NewReader(s string) io.Reader {
	return &reader{start: 0, s: []byte(s)}
}
