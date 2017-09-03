package main

import (
	"golang.org/x/net/html"
	"fmt"
	"gopl/src/ch5/outline2"
	"os"
	"net/http"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err on gettng %s %s", url, err)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("err on parsing html %s %s", url, err)
		return
	}
	title, err := soleTitle(doc)
	if err != nil {
		fmt.Printf("err on getting title %s %s", url, err)
		return
	}
	fmt.Println(title)
}

// soleTitle returns the text of the first non-empty title element
// in doc, and an error if there was no exactly one
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// expected panic
			err = fmt.Errorf("Multitple title")
		default:
			panic(p)
		}
	}()
	// Bail out of recursion if we find more than one non-empty title
	outline2.ForEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no titlte element")
	}
	return title, nil
}
