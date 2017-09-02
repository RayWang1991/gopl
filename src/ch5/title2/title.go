package main

import (
	"net/http"
	"strings"
	"fmt"
	"golang.org/x/net/html"
	"gopl/src/ch5/outline2"
	"os"
)

func main() {
	if len(os.Args[1:]) <= 0 {
		fmt.Fprintf(os.Stderr, "Usage: input url to find its title")
	}
	title(os.Args[1])

}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// check Content-Type is HTML (e.g., "text/html; charset=utt-8"
	ct := resp.Header.Get("Content-Type")
	defer fmt.Println("body Closed")
	defer resp.Body.Close()
	if ct != "text/html" && ! strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing html %s %s", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	outline2.ForEachNode(doc, visitNode, nil)
	return nil
}
