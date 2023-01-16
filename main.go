package main

import (
	"fmt"
	"github.com/ivanspasov99/rss-reader/pkg/meta"
	"github.com/ivanspasov99/rss-reader/pkg/rss"
	"os"
)

func ParseMock(urls []string) []rss.RssItem {
	return []rss.RssItem{
		{
			Title:       "sub1-title",
			Source:      "main-title",
			SourceURL:   "main-link",
			Description: "sub1-desc",
			Link:        "sub1-link",
		},
		{
			Title:       "sub2-title",
			Source:      "main-title",
			SourceURL:   "main-link",
			Description: "sub2-desc",
			Link:        "sub2-link",
		},
	}
}

func main() {
	f, err := os.Create("./rss.json")
	if err != nil {
		fmt.Println(err)
	}

	if err := meta.ParseFeedAsJSON(nil, f, ParseMock); err != nil {
		fmt.Println(err)
		return
	}

	f.Close()
}
