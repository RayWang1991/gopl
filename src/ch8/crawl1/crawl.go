package main

import (
	"fmt"
	"gopl/src/ch5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	workLists := make(chan []string)
	seen := map[string]bool{}
	// send init urls to work lists
	go func() {
		workLists <- os.Args[1:]
	}()

	// receive lists from work lists
	for list := range workLists {
		for _, url := range list {
			if !seen[url] {
				seen[url] = true
				go func(url string) {
					workLists <- crawl(url)
				}(url)
			}
		}
	}
}
