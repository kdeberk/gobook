package main

import (
	"bytes"
	"fmt"
	"strings"
)

func commaRecur(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRecur(s[:n-3]) + "," + s[n-3:]
}

func commaIter(s string) string {
	var buffer bytes.Buffer

	for 0 < len(s) {
		m := len(s) % 3
		if 0 == m {
			m = 3
		}

		buffer.WriteString(s[:m])
		if m < len(s) {
			buffer.WriteRune(',')
		}
		s = s[m:]
	}
	return buffer.String()
}

func commaFloat(s string) string {
	var buffer bytes.Buffer

	if idx := strings.Index(s, "."); 0 < idx {
		buffer.WriteString(commaIter(s[:idx]))
		buffer.WriteRune('.')
		buffer.WriteString(s[idx+1:])
	} else {
		buffer.WriteString(commaIter(s))
	}
	return buffer.String()
}

func main() {
	fmt.Println(commaRecur("12345678"))
	fmt.Println(commaIter("12345678"))
	fmt.Println(commaFloat("12345678.123"))

	fmt.Println(commaRecur("2345678"))
	fmt.Println(commaIter("2345678"))
	fmt.Println(commaFloat("2345678.123"))

	fmt.Println(commaRecur("345678"))
	fmt.Println(commaIter("345678"))
	fmt.Println(commaFloat("345678.123"))
}
