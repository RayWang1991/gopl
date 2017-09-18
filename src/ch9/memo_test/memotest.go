package memo_test

import (
	"net/http"
	"io/ioutil"
	"testing"
	"time"
	"fmt"
	"log"
	"sync"
)

func HTTPGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	out := make(chan string)
	urls := []string{
		"http://www.bongmi.com",
		"http://www.baidu.com",
		"http://www.taobao.com",
		"http://www.baidu.com",
		"http://www.bongmi.com",
		"http://www.12306.cn",
		"http://www.taobao.com",
		"http://www.12306.cn",
	}
	go func() {
		for _, url := range urls {
			out <- url
		}
		close(out)
	}()
	return out
}

// test type
type M interface {
	Get(key string) (value interface{}, err error)
}

func Sequential(t *testing.T, m M) {
	fmt.Printf("Start testing\n")
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("%s %d %s\n", url, len(value.([]byte)), time.Since(start))
	}
}

func Concurrent(t *testing.T, m M) {
	wg := sync.WaitGroup{}
	for url := range incomingURLs() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%s %d %s\n", url, len(value.([]byte)), time.Since(start))
		}(url)
	}
	wg.Wait()
}
