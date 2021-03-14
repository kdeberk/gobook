package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
)

type nodeKind uint

const (
	textNode nodeKind = iota
	elementNode
)

func (n nodeKind) String() string {
	switch n {
	case textNode:
		return "text"
	case elementNode:
		return "element"
	default:
		return "unknown"
	}
}

type node struct {
	kind     nodeKind
	parent   *node
	children []*node
}

func (n node) String() string {
	if 0 == len(n.children) {
		return fmt.Sprintf("%v", n.kind)
	}
	return fmt.Sprintf("%v: %v", n.kind, n.children)
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var tree node

	cur := &tree
	for {
		tok, err := dec.Token()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok.(type) {
		case xml.StartElement:
			var child node
			child.kind = elementNode
			child.parent = cur
			cur.children = append(cur.children, &child)
			cur = &child
		case xml.EndElement:
			cur = cur.parent
		case xml.CharData:
			var child node
			child.kind = textNode
			cur.children = append(cur.children, &child)
		}
	}

	fmt.Println(tree)
}
