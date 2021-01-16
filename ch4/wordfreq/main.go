package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		counts[sc.Text()]++
	}

	fmt.Println("Word counts:")
	for w, c := range counts {
		fmt.Printf("\t%v: %v\n", w, c)
	}
}
