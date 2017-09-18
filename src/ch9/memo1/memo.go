// package memo1 provides a concurrency-unsafe
// memoization of a function of a type Func
package memo

// a Memo caches the result of calling a Func
type Memo struct {
	f     Func
	cache map[string]result
}

type result struct {
	value interface{}
	err   error
}

// Func is the type of fuction to memoize
type Func func(string) (interface{}, error)

func New(f Func) *Memo { return &Memo{f, make(map[string]result)} }

// get cache or call func
// NOTE: not concurrency-safe!
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		// not hit
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
