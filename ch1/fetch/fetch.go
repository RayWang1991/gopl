package main

import (
	"os"
	"net/http"
	"fmt"
	"strings"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("error whlie fetching:%v with url:%v", err, url)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("error while reading:%v with response:%v", err, resp)
			os.Exit(1)
		}
		fmt.Printf("%s", string(b))
	}
}