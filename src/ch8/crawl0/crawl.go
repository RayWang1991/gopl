package crawl0

import (
	"fmt"
	"gopl/src/ch5/links"
	"log"
	"os"
)

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	workList := os.Args[1:]
	seen := map[string]bool{}
	for len(workList) > 0 {
		items := workList
		workList = nil
		for _, url := range items {
			if !seen[url] {
				seen[url] = true
				workList = append(workList, Crawl(url)...)
			}
		}
	}
}
