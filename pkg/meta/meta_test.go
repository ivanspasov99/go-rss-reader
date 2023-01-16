package meta

import (
	"bytes"
	"github.com/ivanspasov99/rss-reader/pkg/rss"
	"io"
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

// Could be refactored ass TDT with various inputs
func TestParseFeedAsJSON(t *testing.T) {
	var buf bytes.Buffer
	w := io.Writer(&buf)

	if err := ParseFeedAsJSON(nil, w, ParseMock); err != nil {
		t.Errorf("ParseFeedAsJSON returned an error: %v", err)
	}

	expectedJSON := "[{\"Title\":\"sub1-title\",\"Source\":\"main-title\",\"SourceURL\":\"main-link\",\"Link\":\"sub1-link\",\"PublishDate\":\"2022-01-03T15:04:05Z\",\"Description\":\"sub1-desc\"},{\"Title\":\"sub2-title\",\"Source\":\"main-title\",\"SourceURL\":\"main-link\",\"Link\":\"sub2-link\",\"PublishDate\":\"2022-01-03T15:04:05Z\",\"Description\":\"sub2-desc\"}]"
	if buf.String() != expectedJSON {
		t.Errorf("Unexpected JSON output. Got: %s, Want: %s", buf.String(), expectedJSON)
	}

}
