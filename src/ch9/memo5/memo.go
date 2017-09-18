// package memo1 provides a concurrency-unsafe
// memoization of a function of a type Func
package memo

import (
	"fmt"
)

type result struct {
	value interface{}
	err   error
}

type Func func(string) (interface{}, error)

type entry struct {
	result
	ready chan struct{}
}

type request struct {
	key  string
	resp chan result
}

type Memo struct {
	requests chan request
}

func (memo *Memo) Get(key string) (interface{}, error) {
	req := request{key, make(chan result)}
	memo.requests <- req
	res := <-req.resp
	return res.value, res.err
}

func New(f Func) *Memo {
	memo := &Memo{make(chan request)}
	go memo.serve(f)
	return memo
}

func (memo *Memo) serve(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if nil == e {
			// not hit
			fmt.Printf("not hit %s\n", req.key)
			// memo is responsible for call f to find out the result
			e = &entry{result{}, make(chan struct{})}
			cache[req.key] = e
			go func(f Func, req request) {
				e.value, e.err = f(req.key)
				close(e.ready)
			}(f, req)
		}
		// wait for the answer to deliver
		go func(req request) {
			<-e.ready
			req.resp <- e.result
		}(req)
	}
}

func (memo *Memo) Close() {
	close(memo.requests)
}
