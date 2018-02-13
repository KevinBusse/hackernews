package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KevinBusse/hackernews"
)

func main() {
	limit := flag.Int("limit", 10, "limit the amount of news, limit=0 means no limit")

	flag.Parse()
	args := flag.Args()

	var items []*hackernews.Item
	var err error

	switch {
	case len(args) > 1 && args[0] == "user":
		fmt.Println("User")
		user, err := hackernews.GetUser(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(user)
		os.Exit(0)
	case len(args) > 0 && args[0] == "new":
		fmt.Println("New Stories")
		items, err = hackernews.GetNewStories(*limit)
	case len(args) > 0 && args[0] == "best":
		fmt.Println("Best Stories")
		items, err = hackernews.GetBestStories(*limit)
	case len(args) > 0 && args[0] == "top":
		fallthrough
	default:
		fmt.Println("Top Stories")
		items, err = hackernews.GetTopStories(*limit)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, item := range items {
		fmt.Println(item)
	}
}
