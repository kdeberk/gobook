package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v", err)
		os.Exit(1)
	}

	printText(doc)
}

func printText(n *html.Node) {
	switch n.Type {
	case html.TextNode:
		fmt.Println(n.Data)
	case html.ElementNode:
		switch n.Data {
		case "script", "style":
			return
		}
		fallthrough
	default:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			printText(c)
		}
	}
}
