package github

import (
	"fmt"
	"time"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

func (i *Issue) String() string {
	return fmt.Sprintf("#%-5d %9.9s %.55s", i.Number, i.User.Login, i.Title)
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
