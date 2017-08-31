package main

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findlinks(url)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findlinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: status %v", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(doc, nil), nil
}

func visit(node *html.Node, links []string) []string {
	if node == nil || node.Type == html.ErrorNode {
		return links
	}
	if node.Type == html.ElementNode {
		if node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		if node.Data == "img" {
			for _, attr := range node.Attr {
				if attr.Key == "src" {
					links = append(links, attr.Val)
				}
			}
		}
	}
	links = visit(node.NextSibling, links)
	links = visit(node.FirstChild, links)
	return links
}