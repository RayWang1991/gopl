package memo_test

import (
	"testing"

	"gopl/src/ch9/memo_test"
	"gopl/src/ch9/memo5"
)

var httpGetBody = memo_test.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memo_test.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memo_test.Concurrent(t, m)
}
