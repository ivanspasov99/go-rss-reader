package main

import (
	"fmt"
	"github.com/ivanspasov99/rss-reader/pkg/meta"
	"github.com/ivanspasov99/rss-reader/pkg/rss"
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
	if err := meta.ParseFeedAsJSON(nil, "./rss.json", ParseMock); err != nil {
		fmt.Println(err)
		return
	}
}
