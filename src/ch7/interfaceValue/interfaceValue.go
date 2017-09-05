package main

import (
	"io"
	"fmt"
	"os"
	"bytes"
)

func main(){
	var w io.Writer
	fmt.Printf("%T\n",w)
	w = os.Stdout
	w.Write([]byte("Hello world!\n"))
	fmt.Printf("%T\n",w)
	w = &bytes.Buffer{}
	fmt.Printf("%T\n",w)
	w = (*os.File)(nil)
	fmt.Printf("%T\n",w)
	var buf *bytes.Buffer
	w = buf
	fmt.Printf("%T\n",w)
	fmt.Printf("buf is nil? %v\n",buf == nil)
	fmt.Printf("w is nil? %v\n",w == nil)
	w = new(bytes.Buffer)
	fmt.Printf("w is nil? %v\n",w == nil)
	w.Write([]byte("I can write!\n"))
}
