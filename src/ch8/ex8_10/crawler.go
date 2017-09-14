package main

import (
	"flag"
	"sync"
	"fmt"
	"log"
	"gopl/src/ch5/links"
	"os"
	"bufio"
)

// crawler is a concurrent web crawler, it extracts urls from the html, and then crawl it if it's not visited
// the -d flag sets the max depth for crawl, the original url is at dept 1
// the default max depth is 3
// the max concurrency for tcp connect is 20
// ex8.10 add cancel feature

//usage -depth default is 3
var depFlag int

var seen map[string]bool = map[string]bool{}

var syn sync.Mutex

type depthList struct {
	list  []string
	depth int
}

var done = make(chan struct{})

func isCanceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func doCrwal(url string, dep int, sig chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if dep > depFlag || isCanceled() {
		return
	}
	sig <- struct{}{}
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
	flag.IntVar(&depFlag, "d", 3, "set the crawl depth(default is 3)")
	flag.Parse()
	buf := make(chan struct{}, 20)
	var wg = sync.WaitGroup{}

	// canceler
	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			switch input.Text() {
			case "q", "Q", "c", "C":
				close(done)
				fmt.Printf("canceled!\n")
				return
			}
		}
	}()

	// init calls
	for _, url := range flag.Args() {
		wg.Add(1)
		go doCrwal(url, 1, buf, &wg)
	}

	wg.Wait()
	panic(nil)
}

func crawl(url string) []string {
	if isCanceled() {
		return nil
	}
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}
