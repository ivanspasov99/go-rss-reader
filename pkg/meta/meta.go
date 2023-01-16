package meta

import (
	"encoding/json"
	"github.com/ivanspasov99/rss-reader/pkg/logging"
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

	printRssItems(items)
	return nil
}

func printRssItems(items []rss.RssItem) {
	for _, item := range items {
		logging.GetLogger().
			Info().
			Str("Title:", item.Title).
			Str("Source:", item.Source).
			Str("Source URL:", item.SourceURL).
			Str("Link:", item.Link).
			Str("Publish date:", item.PublishDate.String()).
			Str("Description:", item.Description)
	}
}
