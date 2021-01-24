package foreach

import "golang.org/x/net/html"

type nodeFn func(*html.Node) bool

func Node(n *html.Node, pre, post nodeFn) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Node(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
