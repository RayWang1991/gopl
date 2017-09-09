package main

import (
	"os"
	"log"
	"strings"
	"fmt"
	"net"
	"bufio"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("usage clockwall (Name=port)+")
	}
	for _, arg := range args {
		nameAndUrl := strings.Split(arg, "=")
		name := nameAndUrl[0]
		url := nameAndUrl[1]
		fmt.Printf("name: %s url: %s\n", name, url)
		c, err := net.Dial("tcp", url)
		if err != nil {
			log.Fatal(err)
		}
		go printTime(c, name)
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() =="q"{
			return
		}
	}
}

func printTime(c net.Conn, n string) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		fmt.Fprintf(os.Stdout, "%s @ %s\n", n, scanner.Text())
	}
}
