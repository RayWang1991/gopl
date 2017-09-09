package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	c, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	tc := c.(*net.TCPConn)
	// stand for read from the connection
	go func() {
		io.Copy(os.Stdout, tc)
		log.Printf("done reading from connetion!\n")
		//tc.CloseRead()
		done <- true
	}()
	io.Copy(tc, os.Stdin)
	log.Printf("done writing to connetion!\n")
	tc.CloseWrite()
	<-done
}
