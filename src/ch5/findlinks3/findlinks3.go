package main

import (
	"log"
	"fmt"
	"gopl.io/ch5/links"
	"os"
)

func bfs(f func(item string) []string, worklist []string) {
	if f == nil {
		log.Fatal("Usage f must not be nil")
	}
	seen := map[string]bool{}
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, node := range items {
			if !seen[node] {
				seen[node] = true
				worklist = append(worklist, f(node)...)
			}
		}
	}
}

// for every unseen item, it will be called once and only once
func crawl(url string) []string {
	fmt.Println(url)
	links, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return links
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Usage enter urls")
	}
	bfs(crawl, os.Args[1:])
}
