package main

import (
	"time"
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"gopl.io/ch4/github"
	"os"
	"log"
)

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json: "total_count"`
	Items      []*Item
}

type Item struct {
	Number   int
	HTMLURL  string `json: "html_url"`
	Title    string
	State    string
	CreateAt time.Time `json: "create_at"`
	User     *User
	Body     string
}

type User struct {
	Login string
	HTML  string `json: html_url`
}

func main() {
	// limit one month
	now := time.Now().Add(-30 * 24 * time.Hour)
	s := now.Format("2006-01-02")
	qOneMon := "updated:>=" + s
	fmt.Println(qOneMon)
	ars := os.Args[1:]
	ars = append(ars, qOneMon)
	result, err := github.SearchIssues(ars)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

// SearchIssues queries the Github issue tracker
func SearchIssurs(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "q=?" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("seach query failed : %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
