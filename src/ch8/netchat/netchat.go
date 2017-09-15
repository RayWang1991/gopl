package main

import (
	"net"
	"os"
	"log"
	"io"
	"fmt"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal(`usage input {port\d+}`)
	}
	conn, err := net.Dial("tcp", "localhost:"+os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	tcpC := conn.(*net.TCPConn) // let it panic if
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // ignore potential io errors
		fmt.Println("done\n")
		done <- struct{}{}
	}()
	mustCopy(tcpC, os.Stdin)
	tcpC.CloseWrite()
	fmt.Println("eof done\n")
	<-done
	fmt.Println("close connect\n")
	tcpC.CloseRead()
}

func mustCopy(writer io.Writer, reader io.Reader) {
	_, err := io.Copy(writer, reader)
	if err != nil {
		log.Fatal(err)
	}
}
