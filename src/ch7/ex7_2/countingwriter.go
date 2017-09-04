package main

import (
	"io"
	"io/ioutil"
	"fmt"
)

type wrapperWriter struct {
	all int64
	io.Writer
}

func (c *wrapperWriter) Write(p []byte) (n int, err error) {
	n, err = c.Writer.Write(p)
	c.all += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newW := wrapperWriter{0, w}
	return &newW, &newW.all
}

func main() {
	var w io.Writer = ioutil.Discard
	u, i := CountingWriter(w)
	fmt.Fprint(u, "123456789")
	fmt.Println(*i)
}
