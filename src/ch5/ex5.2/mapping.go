package main

import (
	"golang.org/x/net/html"
	"os"
	"log"
	"fmt"
)

func main() {
	root, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal("Error on parsing html file: %v", err)
	}
	count := map[string]int{}
	traversePreorder(root, count)
	for k, v := range count {
		fmt.Printf("%-6.6s  #%d\n", k, v)
	}
}

func traversePreorder(n *html.Node, counts map[string]int) {
	if n == nil || n.Type == html.ErrorNode {
		return
	}
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	traversePreorder(n.NextSibling, counts)
	traversePreorder(n.FirstChild, counts)
}
