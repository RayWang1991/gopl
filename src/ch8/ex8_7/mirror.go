// mirror crawls the given url and save the doc in local, using the same dir
// files and links that are not from the original domain are not fetched

package main

import (
	"net/http"
	"io"
	"flag"
	"log"
	"net/url"
	"fmt"
	"os"
	"sync"
	"golang.org/x/net/html"
	"gopl/src/ch5/outline2"
	"path/filepath"
	"strings"
	"bytes"
)

var base *url.URL

var maxDepth int

var maxConcurrency int

var token chan struct{}
var seen map[string]bool = map[string]bool{}
var seensync sync.Mutex
var wg = new(sync.WaitGroup)

func main() {
	// main is responsible for init args
	// max depth
	// base url
	// send the first worker
	// wait until the worker finishes their jobs
	flag.IntVar(&maxDepth, "depth", 3, "the max depth for crawl, default is 3")
	flag.IntVar(&maxConcurrency, "con", 20, "the max concurrent number for crawl, default is 20")
	flag.Parse()

	// init
	token = make(chan struct{}, maxConcurrency)

	if len(flag.Args()) < 1 {
		log.Fatal("usage: input a host url")
	}
	u, err := url.Parse(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid url: %s\n", err)
	}
	base = u

	wg.Add(1)
	go crawl(base.String(), 0)
	wg.Wait()
}

// then visit the valid url, and check the response
// if it is html file, extract the find links
// rewrite the find links with local url
// save the body as local file
func crawl(url string, n int) {
	defer wg.Done()
	if n > maxDepth {
		return
	}
	token <- struct{}{}
	foundLinks, err := visit(url)
	<-token
	if err != nil {
		return
	}
	for _, l := range foundLinks {
		seensync.Lock()
		see := seen[l]
		if !see {
			seen[l] = true
		}
		seensync.Unlock()
		if !see {
			wg.Add(1)
			go crawl(l, n+1)
		}
	}
}

func visit(rawurl string) ([]string, error) {
	// parsed url
	fmt.Printf("%s\n",rawurl)
	url, err := url.Parse(rawurl)
	if err != nil {
		err = fmt.Errorf("invalidurl %s %s\n", rawurl, err)
		return nil, err
	}
	if url.Host != base.Host {
		return nil,fmt.Errorf("not the same host %s\n",url)
	}

	// getting for url
	// max concurrency limit
	req, err := http.Get(rawurl)
	if err != nil {
		return nil, fmt.Errorf("getting %s %s\n", url, err)
	}
	if req.StatusCode != http.StatusOK {
		req.Body.Close()
		return nil, fmt.Errorf("getting %s status %s\n", url, req.Status)
	}

	//
	defer req.Body.Close()
	contentType := req.Header["Content-Type"]
	if strings.Contains(strings.Join(contentType, ","), "text/html") {
		doc, err := html.Parse(req.Body)
		if err != nil {
			return nil, fmt.Errorf("parsing %s as html err %s\n", rawurl, err)
		}
		// extract links
		// find the linked nodes
		// back up the original links
		// rewrite the links
		foundLinks := links(doc)
		b := &bytes.Buffer{}
		html.Render(b, doc) // ignore the err
		save(req, b)
		return foundLinks, nil
	}
	return nil, fmt.Errorf("no more url")
}

// links find the node who has links
// rewrite the links with local ones
// and return the original links (for further crawl)
func links(doc *html.Node) []string {
	res := []string{}
	visit := func(node *html.Node) {
		if node == nil || node.Type == html.ErrorNode {
			return
		}
		if node.Type == html.ElementNode {
			if node.Data == "a" {
				for i, a := range node.Attr {
					if a.Key == "href" {
						u, err := base.Parse(a.Val)
						if err != nil || u.Host != base.Host {
							fmt.Printf("skipping %s", a.Val)
							continue
						}
						res = append(res, u.String())
						u.Scheme = ""
						u.Host = ""
						u.User = nil
						a.Val = u.String()
						node.Attr[i] = a
					}
				}
			}
		}
	}
	outline2.ForEachNode(doc, visit, nil)
	return res
}

// save use the body or req as input
// create file path using the request url
// create file
// if err, print it (deal with it in place)
func save(req *http.Response, body io.Reader) {
	u := req.Request.URL
	pathname := filepath.Join(u.Host, u.Path)
	if filepath.Ext(pathname) == "" {
		pathname = filepath.Join(pathname, "index.html")
	}
	fmt.Printf("file path: %s\n", pathname)
	err := os.MkdirAll(filepath.Dir(pathname), 0777)
	if err != nil {
		fmt.Printf("mkdir %s failed: %s\n", filepath.Dir(pathname), err)
		return
	}
	b := body
	if b == nil {
		b = req.Body
		req.Body.Close()
	}
	f, err := os.Create(pathname)
	if err != nil {
		fmt.Printf("create %s failed: %s\n", pathname, err)
		return
	}
	_, err = io.Copy(f, b)
	if err != nil {
		fmt.Printf("write %s failed: %s\n", f, err)
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Printf("write %s failed: %s\n", f, err)
		return
	}
}
