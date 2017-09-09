package main

import (
	"net"
	"log"
	"io"
	"time"
	"os"
)

func main() {
	port := os.Args[1]
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(con)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("03:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
