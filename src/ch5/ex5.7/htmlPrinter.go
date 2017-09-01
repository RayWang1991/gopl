package main

import (
	"golang.org/x/net/html"
	"fmt"
	"net/http"
	"os"
	"log"
)

var dep = 0

func main() {
	if len(os.Args) < 1 {
		fmt.Printf("Usage: Prog url\n")
		os.Exit(1)
	}
	doc, err := findDoc(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	printForEach(doc, startElement, endElement, singleElemnt)
}

func findDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parse html %s %v", url, err)
	}
	return doc, nil
}

func printForEach(n *html.Node, pre, post, single func(*html.Node)) {
	if n == nil || n.Type == html.ErrorNode {
		return
	}
	if n.FirstChild == nil && single != nil {
		single(n)
		return
	}
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printForEach(c, pre, post, single)
	}
	if post != nil {
		post(n)
	}
}

func singleElemnt(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s%s/>\n", dep*2, "", n.Data, attr(n))
	case html.TextNode:
		fmt.Printf("%*s<%s%s/>\n", dep*2, "", n.Data, attr(n))
	case html.CommentNode:
		fmt.Printf("%*s<%s%s/>\n", dep*2, "", n.Data, attr(n))
	}
}

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s%s>\n", dep*2, "", n.Data, attr(n))
	case html.TextNode:
		fmt.Printf("%*s<%s%s>\n", dep*2, "", n.Data, attr(n))
	case html.CommentNode:
		fmt.Printf("%*s<%s%s>\n", dep*2, "", n.Data, attr(n))
	}
	dep ++
}

func attr(n *html.Node) string {
	if len(n.Attr) == 0 {
		return ""
	}
	s := ""
	for _, a := range n.Attr {
		s += fmt.Sprintf(" %s=%s", a.Key, a.Val)
	}
	return s
}

func endElement(n *html.Node) {
	dep --
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s</%s>\n", dep*2, "", n.Data)
	case html.TextNode:
		fmt.Printf("%*s</%s>\n", dep*2, "", n.Data)
	case html.CommentNode:
		fmt.Printf("%*s</%s>\n", dep*2, "", n.Data)
	}
}
