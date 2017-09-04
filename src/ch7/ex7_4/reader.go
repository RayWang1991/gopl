package main

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"ch5/outline2"
	"fmt"
)

type Reader struct {
	s string
}

func (r *Reader) Read(b []byte) (n int, err error) {
	n = copy(b, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return n, err
}

func main() {
	s := `<html><h1>hahaha</h1></html>`
	doc, err := html.Parse(&Reader{s})
	if err != nil {
		log.Fatal("html parsing err:%s", err)
	}
	outline2.ForEachNode(doc, func(n *html.Node) {
		fmt.Println(n)
	}, nil)
}
