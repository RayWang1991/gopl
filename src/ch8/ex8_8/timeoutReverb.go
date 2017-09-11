package main

import (
	"fmt"
	"time"
	"strings"
	"net"
	"sync"
	"bufio"
	"log"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

// TODO not complete yet
// concurrency is hard
//!+
func handleConn(c net.Conn) {
	tc := c.(*net.TCPConn) // type assert
	tick := time.NewTimer(2 * time.Second)
	input := bufio.NewScanner(tc)
	line := make(chan string)
	var wg = new(sync.WaitGroup)
	go func() {
		for input.Scan() {
			line <- input.Text()
		}
	}()
	for {
		select {
		case <-tick.C:
			tick.Stop() // important!, terminate the tick's goroutine
			fmt.Fprint(tc, "Time out 10s!\n")
			wg.Wait()
			tc.Close()
			return
		case <-line:
			tick.Reset(0)
			wg.Add(1)
			go func() {
				echo(c, input.Text(), 1*time.Second)
				wg.Done()
			}()
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
