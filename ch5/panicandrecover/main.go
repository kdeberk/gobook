package main

import "fmt"

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()

	foo()
}

func foo() {
	panic("on the streets of london")
}
