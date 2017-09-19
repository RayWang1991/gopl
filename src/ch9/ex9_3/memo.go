// package memo1 provides a concurrency-unsafe
// memoization of a function of a type Func

// ex9_3 provides a block solution for cancelling
// that is, block the call to f until it is canceled or return the answer

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
	done chan struct{}
}

type Memo struct {
	requests chan request
}

func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	req := request{key, make(chan result), done}
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
			if req.isCanceled() {
				fmt.Printf("1 canceled %s\n", req.key)
				goto deliever
			}
			e.value, e.err = f(req.key)
			if req.isCanceled() {
				fmt.Printf("2 canceled %s\n", req.key)
				goto deliever
			}
			cache[req.key] = e
			go func(req request) {
				req.resp <- e.result
				close(e.ready)
			}(req)
		}
		// the deliver wait for broadcaster to announce that the mission is canceled or
		// get the result
		// wait for the answer to deliver
	deliever:
		go func(req request) {
			for {
				select {
				case <-e.ready:
					req.resp <- e.result
				default:
					if req.isCanceled() {
						fmt.Printf("3 canceled %s\n", req.key)
						// return empty one
						req.resp <- result{}
						return
					}
				}
			}
		}(req)
	}
}

func (req *request) isCanceled() bool {
	fmt.Printf("isCanceled?\n")
	select {
	case <-req.done:
		fmt.Printf("true!!!!\n")
		return true
	default:
		fmt.Printf("false!!!!\n")
		return false
	}
}

func (memo *Memo) Close() {
	close(memo.requests)
}
