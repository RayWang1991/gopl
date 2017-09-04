package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	n := 0
	for scanner.Scan() {
		n++
	}
	*c += WordCounter(n)
	return n, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int ,error){
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)

	n := 0
	for scanner.Scan() {
		n++
	}
	*c += LineCounter(n)
	return n, nil
}

func main (){
	strs := []string{"123 \n sfafs\n","234 ff \n llasfd","中国 \n123 今天"}
	wc := WordCounter(0)
	lc := LineCounter(0)
	for _,s := range strs{
		b := []byte(s)
		wc.Write(b)
		lc.Write(b)
		fmt.Println(wc)
		fmt.Println(lc)
	}

	wc = 0
	lc = 0
	for _,s := range strs{
		fmt.Fprint(&wc,s)
		fmt.Fprint(&lc,s)
		fmt.Println(wc)
		fmt.Println(lc)
	}
}
