package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), 1*time.Second)
	}
}

func echo(w io.Writer, t string, d time.Duration) {
	fmt.Fprintln(w, "\t", strings.ToUpper(t))
	time.Sleep(d)
	fmt.Fprintln(w, "\t", t)
	time.Sleep(d)
	fmt.Fprintln(w, "\t", strings.ToLower(t))
}
