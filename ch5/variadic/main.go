package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"dberk.nl/gobook/ch5/foreach"
	"dberk.nl/gobook/ch5/prettyprint"
	"golang.org/x/net/html"
)

func main() {
	xs := []int{1, 2, 3, 4, 5}
	fmt.Println(xs, sum(xs...))
	fmt.Println(xs, max(xs[0], xs[1:]...))
	fmt.Println(xs, min(xs[0], xs[1:]...))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(ss, join(", ", ss))

	resp, err := http.Get("https://gopl.io")
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get: %v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse: %v", err)
		os.Exit(1)
	}

	for _, img := range ElementsByTagName(doc, "img") {
		prettyprint.HTML(img)
	}
	for _, header := range ElementsByTagName(doc, "h1", "h2", "h3") {
		prettyprint.HTML(header)
	}
}

func sum(xs ...int) int {
	var s int
	for _, x := range xs {
		s += x
	}
	return s
}

func max(x int, xs ...int) int {
	for _, y := range xs {
		if x < y {
			x = y
		}
	}
	return x
}

func min(x int, xs ...int) int {
	for _, y := range xs {
		if y < x {
			x = y
		}
	}
	return x
}

func join(glue string, ss []string) string {
	var buf bytes.Buffer
	for i, s := range ss {
		buf.WriteString(s)
		if i < len(ss)-1 {
			buf.WriteString(glue)
		}
	}
	return buf.String()
}

func ElementsByTagName(n *html.Node, names ...string) []*html.Node {
	elements := []*html.Node{}

	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, name := range names {
				if name == n.Data {
					elements = append(elements, n)
					break
				}
			}
		}
		return true
	}
	foreach.Node(n, pre, nil)
	return elements
}
