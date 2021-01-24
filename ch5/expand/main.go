package main

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile("\\$[a-zA-Z]+")

func main() {
	m := map[string]string{"what": "Hello", "place": "World"}
	f := func(s string) string { return m[s] }
	s := "$what, $place!"
	fmt.Printf("%s => %s\n", s, expand(s, f))
}

func expand(s string, f func(string) string) string {
	for {
		locs := re.FindStringIndex(s)
		if nil == locs {
			break
		}

		i, j := locs[0], locs[1]
		s = s[:i] + f(s[i+1:j]) + s[j:]
	}
	return s
}
