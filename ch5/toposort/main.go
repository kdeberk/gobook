package main

import (
	"fmt"
	"os"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	// "linear algebra":        {"calculus"},
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "toposort: %v\n", err)
		os.Exit(1)
	}

	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(item string, stack []string) error
	visitAll = func(item string, stack []string) error {
		for _, s := range stack {
			if item == s {
				return fmt.Errorf("cycle found: %#v, %v", stack, item)
			}
		}

		if seen[item] {
			return nil
		}
		seen[item] = true

		for _, c := range m[item] {
			if err := visitAll(c, append(stack, item)); err != nil {
				return err
			}
		}
		order = append(order, item)
		return nil
	}

	for key := range m {
		if err := visitAll(key, make([]string, 0, len(m))); err != nil {
			return nil, err
		}
	}
	return order, nil
}
