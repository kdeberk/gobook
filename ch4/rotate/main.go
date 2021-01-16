package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func rotateSinglePass(s []int, n int) {
	t := make([]int, 0, len(s))
	t = append(t, s[n:]...)
	t = append(t, s[:n]...)

	for i := 0; i < len(s); i++ {
		s[i] = t[i]
	}
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(x)
	rotate(x, 2)
	fmt.Println(x)

	x = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(x)
	rotateSinglePass(x, 2)
	fmt.Println(x)
}
