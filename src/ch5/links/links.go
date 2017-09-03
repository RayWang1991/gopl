package links

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
	"gopl/src/ch5/outline2"
)

// Extract makes an HTTP GET request to the url,
// parses the response as HTML, and returns the
// links in the HTML document
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("getting:%s %s", url, err)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing HTML %s %s", url, err)
	}
	links := []string{}
	var visitNode = func(n *html.Node) {
		if n == nil || n.Type == html.ErrorNode {
			return
		}
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
	}
	outline2.ForEachNode(doc, visitNode, nil)
	return links, nil
}
