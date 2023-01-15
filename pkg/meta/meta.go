package meta

import (
	"encoding/json"
	"fmt"
	"github.com/ivanspasov99/rss-reader/pkg/rss"
	"os"
)

type RSSParse func([]string) []rss.RssItem

// ParseFeedAsJSON parses feed, stores it in a file as JSON and prints it on the stdout
func ParseFeedAsJSON(urls []string, filepath string, parse RSSParse) error {
	items := parse(urls)

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err := enc.Encode(items); err != nil {
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
