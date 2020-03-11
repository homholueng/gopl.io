package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/homholueng/gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	format := "#%-5d %9.9s %v %.55s\n"

	now := time.Now()
	oneMonthBefore := now.AddDate(0, -1, 0)
	oneYearBefore := now.AddDate(-1, 0, 0)

	inMonth := make([]*github.Issue, 0)
	inYear := make([]*github.Issue, 0)
	beforeYear := make([]*github.Issue, 0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(oneMonthBefore):
			inMonth = append(inMonth, item)
		case item.CreatedAt.After(oneYearBefore) && item.CreatedAt.Before(oneMonthBefore):
			inYear = append(inYear, item)
		case item.CreatedAt.Before(oneYearBefore):
			beforeYear = append(beforeYear, item)
		}
	}
	if len(inMonth) > 0 {
		fmt.Printf("\nIn month:\n")
		for _, item := range inMonth {
			fmt.Printf(format, item.Number, item.User.Login, item.CreatedAt, item.Title)
		}
	}
	if len(inYear) > 0 {
		fmt.Printf("\nIn year:\n")
		for _, item := range inYear {
			fmt.Printf(format, item.Number, item.User.Login, item.CreatedAt, item.Title)
		}
	}
	if len(beforeYear) > 0 {
		fmt.Printf("\nBefore year:\n")
		for _, item := range beforeYear {
			fmt.Printf(format, item.Number, item.User.Login, item.CreatedAt, item.Title)
		}
	}
}
