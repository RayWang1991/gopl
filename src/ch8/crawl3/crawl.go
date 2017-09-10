package main

import (
	"gopl/src/ch8/crawl0"
	"os"
)

func main() {
	seen := map[string]bool{}
	foundLinks := make(chan []string)
	unseenLinks := make(chan string)
	go func() {
		foundLinks <- os.Args[1:]
	}()

	// create 20 worker
	// worker receive url from the unseen links, and crawl new links
	// and send them to the found links
	for i := 0; i < 20; i++ {
		go func() {
			for url := range unseenLinks {
				newOne := crawl0.Crawl(url)
				// new goroutine to break the dead lock
				go func() {
					foundLinks <- newOne
				}()
			}
		}()
	}

	// the main goroutine de-duplicate the found links,
	// send the new ones to the unseen links
	for list := range foundLinks {
		for _, url := range list {
			if !seen[url] {
				seen[url] = true
				unseenLinks <- url
			}
		}
	}
}
