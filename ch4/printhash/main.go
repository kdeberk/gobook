package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var bits = flag.Int("bits", 256, "bitsize")

func main() {
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	switch *bits {
	case 256:
		fmt.Println(sha256.Sum256([]byte(input)))
	case 384:
		fmt.Println(sha512.Sum384([]byte(input)))
	case 512:
		fmt.Println(sha512.Sum512([]byte(input)))
	default:
		fmt.Fprintf(os.Stderr, "Unknown bitsize %v", *bits)
		os.Exit(1)
	}
}
