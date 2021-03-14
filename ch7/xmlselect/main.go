package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement

	searchItems := parseSearchItems(os.Args[1:])
	fmt.Println(searchItems)

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, searchItems) {
				ns := make([]string, 0, len(stack))
				for _, el := range stack {
					ns = append(ns, el.Name.Local)
				}
				fmt.Printf("%s: %s\n", strings.Join(ns, " "), tok)
			}
		}
	}
}

type searchItem struct {
	name, id, class string
}

// parseSearchItems parses CSS-style search strings
func parseSearchItems(ss []string) []searchItem {
	si := make([]searchItem, len(ss))
	for i, s := range ss {
		prev := len(s)
		for j := len(s) - 1; 0 <= j; j-- {
			switch s[j] {
			case '.':
				si[i].class = s[j+1 : prev]
				prev = j
			case '#':
				si[i].id = s[j+1 : prev]
				prev = j
			}
		}
		si[i].name = s[0:prev]
	}
	return si
}

// containsAll reports whether els contains the elements of y, in order.
func containsAll(els []xml.StartElement, is []searchItem) bool {
	for len(is) <= len(els) {
		if len(is) == 0 {
			return true
		}
		if elementMatches(els[0], is[0]) {
			is = is[1:]
		}
		els = els[1:]
	}
	return false
}

func elementMatches(el xml.StartElement, i searchItem) bool {
	if el.Name.Local != i.name {
		return false
	}

	if i.id != "" {
		if attr, ok := findAttr(el, "id"); !ok || attr.Value != i.id {
			return false
		}
	}

	if i.class != "" {
		if attr, ok := findAttr(el, "class"); !ok || attr.Value != i.class {
			return false
		}
	}

	return true
}

func findAttr(el xml.StartElement, n string) (attr xml.Attr, ok bool) {
	for _, attr := range el.Attr {
		if attr.Name.Local == n {
			return attr, true
		}
	}
	return
}
