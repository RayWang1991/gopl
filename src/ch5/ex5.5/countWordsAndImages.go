package main

import (
	"net/http"
	"golang.org/x/net/html"
	"fmt"
	"bufio"
	"strings"
	"os"
	"log"
	"gopl/src/ch5/outline2"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "usage: PROG URL")
	}
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("words: %d images: %d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML %s %v", url, err)
		return
	}
	words, images = countWordsAndImages(doc)
	outline2.Outline2(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil || n.Type == html.ErrorNode {
		return //0,0
	}
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images = 1
		}
	} else if n.Type == html.TextNode {
		words = countWord(n.Data)
	}
	delW1, delI1 := countWordsAndImages(n.FirstChild)
	delW2, delI2 := countWordsAndImages(n.NextSibling)
	words += delW1 + delW2
	images += delI1 + delI2
	return
}

func countWord(s string) (n int) {
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return
}
