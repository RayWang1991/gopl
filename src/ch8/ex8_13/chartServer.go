package main

import (
	"net"
	"bufio"
	"log"
	"fmt"
	"sort"
	"strings"
	"time"
)

// ex8_13, disconnect the idle clients
// ex8_12, register user names
// chart server is a chart server...
// clients connect to it, send messages to broadcaster
// then the broadcaster send messages to all clients

type clientReader chan string

type clientInfo struct {
	c    clientReader
	name string
}

var (
	enter = make(chan *clientInfo)
	leave = make(chan clientReader)
	msg   = make(chan string)
)

func broadcast() {
	// clients set
	clients := map[clientReader]string{}

	for {
		select {
		case c := <-enter:
			// update all names
			clients[c.c] = c.name
			allNames := make([]string, 0, len(clients))
			for _, v := range clients {
				allNames = append(allNames, v)
			}
			sort.Strings(allNames)
			// broadcast to all member
			inMsg := fmt.Sprintf("%s came, welcome", c.name) + " current member:" + strings.Join(allNames, ",")
			for c := range clients {
				c <- inMsg
			}
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

	// input name
	fmt.Fprintln(conn, "welcome, please enter your name:")
	input := bufio.NewScanner(conn)
	var who string
	for input.Scan() {
		who = input.Text()
		break
	}

	// set tup reader
	c := clientReader(make(chan string))
	go readClient(conn, c, who)

	// enter chart room
	enter <- &clientInfo{c, who}
	fmt.Printf("%s has came\n", who)

	// role as client writer
	prefix := who + ":"
	cInput := make(chan string)

	go func(cIn chan<- string) {
		for input.Scan() {
			cInput <- input.Text()
		}
		close(cInput)
	}(cInput)

	// t in 10 s
	t := time.NewTimer(10 * time.Second)

loop:
	for {
		select {
		case <-t.C:
			// leave
			fmt.Printf("%s leave by Time Out\n",who)
			break loop
		case text, ok := <-cInput:
			if !ok {
				fmt.Printf("%s leave by EOF\n",who)
				break loop
			}
			t.Reset(10 * time.Second)
			msg <- prefix + text
		}
	}
	t.Stop()
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
