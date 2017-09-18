// package memo1 provides a concurrency-unsafe
// memoization of a function of a type Func
package memo

import (
	"sync"
	"fmt"
)

// a Memo caches the result of calling a Func
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string] *entry
}

// entry contains result
// and the ready channel indicating that whether
// the result is ready for reading
type entry struct{
	res result
	ready chan struct{}
}

type result struct {
	value interface{}
	err   error
}

// Func is the type of function to memoize
type Func func(string) (interface{}, error)

func New(f Func) *Memo { return &Memo{f, sync.Mutex{}, make(map[string] *entry)} }

// get cache or call func
// NOTE: Get is concurrency-safe, once an entry is ready, just return the value
// otherwise all calling goroutines wait for the answer given by the first goroutine.
func (memo *Memo) Get(key string) (interface{}, error) {
	// read is ok
	memo.mu.Lock()
	e, ok := memo.cache[key]
	if !ok {
		fmt.Printf("not hit with %s",key)
		// not hit
		// the first goroutine is responsible for get the result
		e = &entry{ready:make (chan struct{})}
		e.res.value, e.res.err = memo.f(key)
		memo.cache[key] = e
		close(e.ready)
		memo.mu.Unlock()
	} else {
		fmt.Printf("hit with %s",key)
		memo.mu.Unlock()
		// Unlock first, all it needs to do is wait for the answer
		<- e.ready
	}
	return e.res.value, e.res.err
}
