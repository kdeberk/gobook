package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, ns := range counts {
		total := 0
		filenames := []string{}

		for filename, n := range ns {
			total += n
			filenames = append(filenames, filename)
		}

		if 1 < total {
			fmt.Printf("%d\t%s\t%v\n", total, line, filenames)
		}
	}
}

func countLines(f *os.File, name string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t := input.Text()
		if _, ok := counts[t]; !ok {
			counts[t] = make(map[string]int)
		}
		counts[input.Text()][name]++
	}
}
