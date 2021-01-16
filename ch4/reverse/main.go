package main

import "fmt"

func reverse_slice(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse_arrayptr(aptr [10]byte) {
	for i, j := 0, len(aptr)-1; i < j; i, j = i+1, j-1 {
		aptr[i], aptr[j] = aptr[j], aptr[i]
	}
}

func reverse_unicode(bs []byte) {
	reverse_slice(bs)
	for i := len(bs) - 1; 0 <= i; {
		switch {
		case 0b00000000 == 0b10000000&bs[i]:
			// 1 byte
			i -= 1
		case 0b11000000 == 0b11100000&bs[i]:
			// 2 bytes
			reverse_slice(bs[i-1 : i+1])
			i -= 2
		case 0b11100000 == 0b11110000&bs[i]:
			// 3 bytes
			reverse_slice(bs[i-2 : i+1])
			i -= 3
		case 0b11110000 == 0b11111000&bs[i]:
			// 4 bytes
			reverse_slice(bs[i-3 : i+1])
			i -= 4
		}
	}
}

func reverse(s string, rFn func([]byte)) string {
	bs := []byte(s)
	rFn(bs)
	return string(bs)
}

func main() {
	var s string

	s = "Hello, World!"
	fmt.Println(s, reverse(s, reverse_slice))
	fmt.Println(s, reverse(s, reverse_unicode))

	s = "Hello, 世界한글!"
	fmt.Println(s, reverse(s, reverse_slice))
	fmt.Println(s, reverse(s, reverse_unicode))
}
