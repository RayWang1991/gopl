package main

import (
	"os"
	"golang.org/x/net/html"
	"log"
	"fmt"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal("Error on parsing %v\n", err)
	}
	for i, s := range visit(doc, []string{}) {
		fmt.Printf("Link #%-4d %s\n", i, s)
	}
}

func visit(n *html.Node, links []string) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(n.FirstChild, links)
	links = visit(n.NextSibling, links)
	return links
}
