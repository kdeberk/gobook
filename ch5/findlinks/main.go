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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	getLink := func(links []string, el string, as []html.Attribute) []string {
		for _, a := range as {
			if a.Key == el {
				return append(links, a.Val)
			}
		}
		return links
	}

	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "style":
			links = getLink(links, "href", n.Attr)
		case "img", "script":
			links = getLink(links, "src", n.Attr)
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}
