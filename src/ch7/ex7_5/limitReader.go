package main

import (
	"io"
	"strings"
	"fmt"
)

type limitReader struct {
	r   io.Reader
	lim int64
}

func (l *limitReader) Read(p []byte) (res int, err error) {
	m := int64(len(p))
	if l.lim < m {
		m = l.lim
	}
	res, err = l.r.Read(p[:m])
	l.lim -= int64(res)
	if l.lim <= 0 {
		err = io.EOF
	}
	return
}

func main() {
	var b = make([]byte, 4)
	l := &limitReader{strings.NewReader("1234567"), 5}
	n, err := l.Read(b)
	fmt.Printf("n: %d, err:%s\n", n, err)
	fmt.Println(string(b))
}
