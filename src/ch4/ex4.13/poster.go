package main

import (
	"net/http"
	"net/url"
	"fmt"
	"encoding/json"
	"os"
)

const omdbURL = "https://omdbapi.com/"

type movie struct {
	Title   string
	Year    int
	Runtime int
	Actors  []string
	Poster  string
	Website string
}

func main() {
	if len(os.Args) <= 1 {
		return
	}
	arg := os.Args[1]
	movie, err := getMovie(arg)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("%#v", movie)
}

func getMovie(title string) (*movie, error) {
	title = url.QueryEscape(title)
	fmt.Println(omdbURL + "?t=" + title)
	resp, err := http.Get(omdbURL + "?t=" + title)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("%v", resp.Status)
	}
	var res movie
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&res); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &res, nil
}
