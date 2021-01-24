package prettyprint

import (
	"fmt"

	"dberk.nl/gobook/ch5/foreach"
	"golang.org/x/net/html"
)

func HTML(n *html.Node) {
	var depth int

	pre := func(n *html.Node) bool {
		switch n.Type {
		case html.ElementNode:
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			depth++

			for _, a := range n.Attr {
				fmt.Printf(` %s="%s"`, a.Key, a.Val)
			}

			switch n.FirstChild {
			case nil:
				fmt.Printf("/>\n")
			default:
				fmt.Printf(">\n")
			}
		case html.TextNode:
			fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		}
		return true
	}
	post := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			depth--
			if nil != n.FirstChild {
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		}
		return true
	}

	foreach.Node(n, pre, post)
}
