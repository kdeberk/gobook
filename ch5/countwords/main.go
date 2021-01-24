package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if 1 == len(os.Args) {
		fmt.Printf("Usage: ./countwords <url>")
		return
	}

	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countwords: %v", err)
		}
		fmt.Printf("%v\n\twords: %d, images: %d\n", url, words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %v", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	switch n.Type {
	case html.TextNode:
		words += len(strings.Split(n.Data, " "))
	case html.ElementNode:
		switch n.Data {
		case "style", "script":
			return
		case "img":
			images++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ws, is := countWordsAndImages(c)
		words += ws
		images += is
	}
	return
}
