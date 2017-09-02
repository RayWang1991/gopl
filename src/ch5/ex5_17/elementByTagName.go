package main

import (
	"golang.org/x/net/html"
	"gopl/src/ch5/outline2"
	"os"
	"net/http"
	"fmt"
)

func main() {
	url := os.Args[1]
	names := os.Args[2:]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getting %s %s", url, err)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing HTML %s %s", url, err)
	}
	nodes := ElementByTagName(doc, names...)
	for _, n := range nodes {
		fmt.Printf("node %s\n", n.Data)
	}
}

func ElementByTagName(doc *html.Node, names ...string) []*html.Node {
	set := map[string]bool{}
	for _, n := range names {
		set[n] = true
	}
	res := []*html.Node{}
	f := func(node *html.Node) {
		if set[node.Data] {
			res = append(res, node)
		}
	}
	outline2.ForEachNode(doc, f, nil)
	return res
}
