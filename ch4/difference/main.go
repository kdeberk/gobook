package main

import (
	"crypto/sha256"
	"fmt"
)

func popDifference(a, b [32]byte) int {
	c := 0
	for i := 0; i < 32; i++ {
		if a[i] != b[i] {
			c++
		}
	}
	return c
}

func main() {
	a := sha256.Sum256([]byte("foo"))
	b := sha256.Sum256([]byte("bar"))

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(popDifference(a, b))
}
