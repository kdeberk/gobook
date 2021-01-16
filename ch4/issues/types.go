package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"dberk.nl/gobook/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	thisMonth := []*github.Issue{}
	thisYear := []*github.Issue{}
	older := []*github.Issue{}

	t := time.Now()
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(t.AddDate(0, -1, 0)):
			thisMonth = append(thisMonth, item)
		case item.CreatedAt.After(t.AddDate(-1, 0, 0)):
			thisYear = append(thisYear, item)
		default:
			older = append(older, item)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("%d this month\n", len(thisMonth))
	for _, item := range thisMonth {
		fmt.Println(item)
	}
	fmt.Printf("%d this year\n", len(thisYear))
	for _, item := range thisYear {
		fmt.Println(item)
	}
	fmt.Printf("%d older\n", len(older))
	for _, item := range older {
		fmt.Println(item)
	}
}
