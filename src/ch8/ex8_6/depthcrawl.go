package main

import (
	"flag"
	"sync"
	"fmt"
	"log"
	"gopl/src/ch5/links"
)

//usage -depth default is 3
var depFlag int

var seen map[string]bool = map[string]bool{}

var syn sync.Mutex

type depthList struct {
	list  []string
	depth int
}

func doCrwal(url string, dep int, sig chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if dep >= depFlag {
		return
	}
	sig <- struct{}{}
	// n+1 means currently is n + 1 depth for the url list
	linkUrls := crawl(url)
	<-sig

	// do dfs
	for _, l := range linkUrls {
		syn.Lock()
		_, ok := seen[l]
		if !ok {
			seen[l] = true
		}
		syn.Unlock()
		if !ok {
			wg.Add(1)
			go doCrwal(l, dep+1, sig, wg)
		}
	}
}

// a concurrent crawler for at most 20 worker
func main() {
	flag.IntVar(&depFlag, "depth", 3, "set the crawl depth(default is 3)")
	flag.Parse()
	buf := make(chan struct{}, 20)
	var wg = sync.WaitGroup{}
	// init calls

	for _, url := range flag.Args() {
		wg.Add(1)
		go doCrwal(url, 0, buf, &wg)
	}

	wg.Wait()
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}
