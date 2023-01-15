package meta

import (
	"encoding/json"
	"fmt"
	"github.com/ivanspasov99/rss-reader/pkg/rss"
	"os"
	"reflect"
	"testing"
	"time"
)

func ParseMock(urls []string) []rss.RssItem {
	return []rss.RssItem{
		{
			Title:       "sub1-title",
			Source:      "main-title",
			SourceURL:   "main-link",
			Description: "sub1-desc",
			PublishDate: time.Date(2022, 1, 3, 15, 4, 5, 0, time.UTC),
			Link:        "sub1-link",
		},
		{
			Title:       "sub2-title",
			Source:      "main-title",
			SourceURL:   "main-link",
			Description: "sub2-desc",
			PublishDate: time.Date(2022, 1, 3, 15, 4, 5, 0, time.UTC),
			Link:        "sub2-link",
		},
	}
}

func TestParseFeedAsJSON(t *testing.T) {
	tempFile, err := os.CreateTemp("", "rssitems.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	if err := ParseFeedAsJSON(nil, tempFile.Name(), ParseMock); err != nil {
		t.Fatal("Unexpected error:", err)
	}

	if _, err := os.Stat(tempFile.Name()); os.IsNotExist(err) {
		t.Fatal("Expected file to be created but it does not exist")
	}

	var items []rss.RssItem
	if err := json.NewDecoder(tempFile).Decode(&items); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	if !reflect.DeepEqual(items, ParseMock(nil)) {
		t.Fatal("Expected items are not the same as in the file")
	}

}
