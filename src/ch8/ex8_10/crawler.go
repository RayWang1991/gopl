package main

import (
	"flag"
	"sync"
	"fmt"
	"log"
	"os"
	"bufio"
	"net/http"
	"golang.org/x/net/html"
	"gopl/src/ch5/outline2"
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
	fmt.Println(url)
	list, err := extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func extract(url string) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if nil != err {
		return nil, err
	}
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getting:%s %s", url, err)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing HTML %s %s", url, err)
	}
	links := []string{}
	var visitNode = func(n *html.Node) {
		if n == nil || n.Type == html.ErrorNode {
			return
		}
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
	}
	outline2.ForEachNode(doc, visitNode, nil)
	return links, nil
}
