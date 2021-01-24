package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
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
