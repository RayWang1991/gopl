package outline2

import (
	"golang.org/x/net/html"
	"fmt"
)

func Outline2(node *html.Node) {
	forEachNode(node, startElement, endElement)
}

func forEachNode(node *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(node)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(node)
	}
}

var depth int = 0

func startElement(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
	}
	depth++
}

func endElement(node *html.Node) {
	depth--
	if node.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
	}
}
