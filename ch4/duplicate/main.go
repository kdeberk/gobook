package main

import "fmt"

func main() {
	ss := []string{"foo", "foo", "bar", "bar", "foo", "baz", "baz", "baz", "baz", "foo"}
	fmt.Println(ss)
	fmt.Println(removeDuplicates(ss))
}

func removeDuplicates(s []string) []string {
	i, j := 0, 1
	for j < len(s) {
		switch {
		case s[i] == s[j]:
			j++
		default:
			s[i+1] = s[j]
			i++
			j++
		}
	}
	return s[:i+1]
}
