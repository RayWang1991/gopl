package memo_test

import (
	"testing"

	"gopl/src/ch9/ex9_3"
	"log"
	"fmt"
	"time"
	"sync"
	"net/http"
	"io/ioutil"
)

var httpGetBody = HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	Concurrent(t, m)
}

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
	Get(key string, done chan struct{}) (value interface{}, err error)
}

func Sequential(t *testing.T, m M) {
	fmt.Printf("Start testing\n")
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, make(chan struct{}))
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
			done := make(chan struct{})
			go func() {
				time.Sleep(1000 * time.Millisecond)
				close(done)
			}()
			value, err := m.Get(url, done)
			if err != nil {
				log.Println(err)
				return
			}
			if value == nil {
				log.Println("value is nil, cancel")
				return
			}
			fmt.Printf("%s %d %s\n", url, len(value.([]byte)), time.Since(start))
		}(url)
	}
	wg.Wait()
}
