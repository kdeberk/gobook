package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"os"
	"strings"
	"time"
)

func main() {
	var ps []**string
	for _, a := range os.Args[1:] {
		p, err := readClock(a)
		if err != nil {
			log.Fatal(err)
		}
		ps = append(ps, p)
	}

	for {
		var ss []string
		for _, p := range ps {
			ss = append(ss, **p)
		}

		fmt.Printf("| %v |\n", strings.Join(ss, " | "))
		time.Sleep(1 * time.Second)
	}
}

func readClock(url string) (**string, error) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		return nil, err
	}
	reader := textproto.NewReader(bufio.NewReader(conn))

	var l string
	var p *string = &l

	go func() {
		defer conn.Close()

		for {
			l, err := reader.ReadLine()
			if err != nil {
				break
			}
			p = &l
		}
	}()
	return &p, nil
}
