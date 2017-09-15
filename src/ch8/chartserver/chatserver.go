package main

import (
	"net"
	"bufio"
	"log"
	"fmt"
)

// chart server is a chart server...
// clients connect to it, send messages to broadcaster
// then the broadcaster send messages to all clients

type clientReader chan string

var (
	enter = make(chan clientReader)
	leave = make(chan clientReader)
	msg   = make(chan string)
)

// TODO return if all clients are gone ?
func broadcast() {
	// clients set
	clients := map[clientReader]bool{}
	for {
		select {
		case c := <-enter:
			clients[c] = true
		case c := <-leave:
			delete(clients, c)
		case m := <-msg:
			for c := range clients {
				c <- m
			}
		}
	}
}

func handleConn(conn net.Conn) {
	// initiation for conn
	who := conn.RemoteAddr().String()
	c := clientReader(make(chan string))
	go readClient(conn, c, who)
	enter <- c
	fmt.Printf("%s has came\n", who)
	c <- "you are" + who
	msg <- who + "have came, welcome!"
	// role as client writer
	prefix := who + ":"
	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg <- prefix + input.Text()
	}
	msg <- prefix + "has left"
	leave <- c
	close(c)
	conn.Close()
}

func readClient(conn net.Conn, c clientReader, who string) {
	prefix := who + ":"
	for text := range c {
		fmt.Fprintln(conn, prefix+text)
	}
}

func main() {
	go broadcast()
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("failed to listen", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("connect to %s %s", listener.Addr().String(), err)
			continue
		}
		fmt.Printf("connection to %s\n", conn.RemoteAddr().String())
		go handleConn(conn)
	}
}
