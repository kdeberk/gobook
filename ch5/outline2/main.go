package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"dberk.nl/gobook/ch5/foreach"
	"dberk.nl/gobook/ch5/prettyprint"
	"golang.org/x/net/html"
)

var (
	searchID = flag.String("id", "", "Element with ID to search")
	pretty   = flag.Bool("pretty", false, "Whether to prettyprint HTML")
	url      = flag.String("url", "", "URL to visit")
)

func main() {
	flag.Parse()

	if "" == *url {
		flag.Usage()
		os.Exit(0)
	}

	resp, err := http.Get(*url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline2: %v", err)
		os.Exit(1)
	}

	switch {
	case *pretty:
		prettyprint.HTML(doc)
	case *searchID != "":
		searchHTMLForID(doc, *searchID)
	default:
		flag.Usage()
	}
}

func searchHTMLForID(n *html.Node, id string) {
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					prettyprint.HTML(n)
					return false
				}
			}
		}
		return true
	}

	foreach.Node(n, pre, nil)
}
