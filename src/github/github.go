//package github provide API of Github issue
package github

import  (
    "time"
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
    TotalCount int `json:"html_url"`
    Items []*Issue
}

type Issue struct {
    Number int
    HTMLURL string `json:"html_url"`
    Title string
    State string
    User *User
    CreatedAt time.Time `json:"created_at"`
    Body string
}

type User struct {
    Login string
    HTMLURL string `json:html_url`
}

func SearchIssues(terms []string)(*IssuesSearchResult, error) {
    q := url.QueryEscape(strings.Join(terms, " "))
    resp, err := http.Get(IssuesURL + "?q=" +q)
    if err != nil {
        return nil, err
    }


    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("Search query failed: %s", resp.Status)
    }

    var result IssuesSearchResult
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil, err
    }

    resp.Body.Close()
    return &result, nil
}
