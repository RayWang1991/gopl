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
	"html/template"
)

const IssueURL = "https://api.github.com/search/issues"

const temp1 = `{{.TotalCount}} issues:
{{range.Items}}---------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

type IssueSearchResult struct {
	TotalCount int `json: "total_count"`
	Items      []*Item
}

type Item struct {
	Number    int
	HTMLURL   string `json: "html_url"`
	Title     string
	State     string
	CreatedAt time.Time `json: "create_at"`
	User      *User
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json: html_url`
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
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
	/*
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	*/
	// using template1
	/*
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(temp1)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
	*/
	// output html
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
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

var issueList = template.Must(template.New(`issuelist`).
	Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range.Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></th>
  <td>{{.State}}</th>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></th>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></th>
</tr>
{{end}}
</table>
`))
