package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	sortedResult := make(map[string][]*github.Issue)
	now := time.Now()

	for _, item := range result.Items {
		delta := now.Sub(item.CreatedAt).Hours() / 24
		switch {
		case delta < 30:
			sortedResult["before30"] = append(sortedResult["before30"], item)
		case delta < 365:
			sortedResult["before365"] = append(sortedResult["before30"], item)
		case delta >= 365:
			sortedResult["after365"] = append(sortedResult["before30"], item)
		}
	}

	for key, items := range sortedResult {
		fmt.Printf("Created - %s\n", key)
		for _, item := range items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
