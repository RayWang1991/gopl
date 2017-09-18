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

// Func is the type of function to memoize
type Func func(string) (interface{}, error)

func New(f Func) *Memo { return &Memo{f, sync.Mutex{}, make(map[string]result)} }

// get cache or call func
// NOTE: Get is concurrency-safe, using two critical-region.
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		// not hit
		// call memo.f () concurrently is ok, but there is a chance for several goroutines to call the redundantly
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
