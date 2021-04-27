package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	name string
	c    chan<- string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages.
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients.
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming messages to all clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.c <- msg:
					continue
				default:
					// If there is no recv on the other side of the channel then skip sending the
					// message.
					continue
				}
			}
		case cli := <-entering:
			if 0 == len(clients) {
				cli.c <- "You're the only one in the channel"
			} else {
				var names []string
				for c := range clients {
					names = append(names, c.name)
				}
				cli.c <- "Present: " + strings.Join(names, ", ")
			}
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.c)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 1) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"

	cli := client{name: who, c: ch}
	entering <- cli

	recvd := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.After(5 * time.Minute):
				conn.Close()
				return
			case <-recvd:
				continue
			}
		}
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		recvd <- struct{}{} // Reset idle timer.
	}
	// NOTE: ignoring potential errors errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOT: ignoring network errors
	}
}
