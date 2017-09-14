package main

import (
	"net/http"
	"fmt"
	"os"
	"runtime"
	"sync"
)

type info struct {
	header http.Header
	host   string
	err    error
}

var done = make(chan struct{})

var wg = sync.WaitGroup{}

// mirrorReq request a list of urls in parallel, return the fastest one, and cancel the others
func main() {
	// canceler
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Printf("usage: enter urls\n")
		return
	}

	res := make(chan *info , len(urls))
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			head, host, err := request(url)
			res <- &info{head, host, err}
			wg.Done()
		}(url)
	}
	info := <-res
	// cancel the others, and wait they are all closed
	close(done)
	fmt.Printf("cancel\n")
	wg.Wait()
	fmt.Printf("host:%s error:%s\n", info.host, info.err)
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	fmt.Printf("%s", buf)
}

func request(hostname string) (header http.Header, host string, err error) {
	req, err := http.NewRequest("GET", hostname, nil)
	if err != nil {
		return nil, "", err
	}
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("getting:%s %s", req, err)
	}
	resp.Body.Close()
	return resp.Header, resp.Request.URL.String(), nil
}
