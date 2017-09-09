package main

import (
	"net"
	"os"
	"log"
	"io"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal(`usage input {port\d+}`)
	}
	conn, err := net.Dial("tcp", "localhost:"+os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(writer io.Writer, reader io.Reader) {
	_, err := io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}
}
