package ex5_12

import (
	"fmt"
	"golang.org/x/net/html"
)

func Outline2(node *html.Node) {
	start, end := outline()
	forEachNode(node, start, end)
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

func outline() (start, end func(node *html.Node)) {
	var depth int
	return func(node *html.Node) {
		if node.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
		}
		depth++
	}, func(node *html.Node) {
		depth--
		if node.Type == html.ElementNode {
			fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
		}
	}
}
