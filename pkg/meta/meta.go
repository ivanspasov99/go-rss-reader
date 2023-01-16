package meta

import (
	"encoding/json"
	"fmt"
	"github.com/ivanspasov99/rss-reader/pkg/rss"
	"io"
)

type RSSParse func([]string) []rss.RssItem

// ParseFeedAsJSON parses feed based on urls used for the parse function, writes it in a stream as JSON and prints it on the stdout
func ParseFeedAsJSON(urls []string, w io.Writer, parse RSSParse) error {
	items := parse(urls)

	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		return err
	}

	for _, item := range items {
		// Log in JSON
		fmt.Println("Title:", item.Title)
		fmt.Println("Source:", item.Source)
		fmt.Println("Source URL:", item.SourceURL)
		fmt.Println("Link:", item.Link)
		fmt.Println("Publish date:", item.PublishDate)
		fmt.Println("Description:", item.Description)
		fmt.Println()
	}
	return nil
}
