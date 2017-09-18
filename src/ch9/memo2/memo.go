// package memo1 provides a concurrency-unsafe
// memoization of a function of a type Func
package memo

import "sync"

// a Memo caches the result of calling a Func
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

type result struct {
	value interface{}
	err   error
}

// Func is the type of fuction to memoize
type Func func(string) (interface{}, error)

func New(f Func) *Memo { return &Memo{f, sync.Mutex{}, make(map[string]result)} }

// get cache or call func
// NOTE: Get is concurrency-safe.
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		// not hit
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
