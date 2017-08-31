package main

import (
	"golang.org/x/net/html"
	"fmt"
	"os"
	"log"
)

func main() {
	root, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal("Error on parsing html %v", err)
	}
	traversePreoder(root)
}

func traversePreoder(node *html.Node) {
	if node == nil || node.Type == html.ErrorNode {
		return
	}
	if node.Type == html.TextNode {
		fmt.Printf("%q\n", node.Data)
	}
	traversePreoder(node.NextSibling)
	traversePreoder(node.FirstChild)
}
