package rss

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var testProcessUrl = []struct {
	name              string
	xmlInputFilePath  string
	expectedRssItems  map[string]RssItem
	expectedItemsSize int
}{
	{
		"Test should parse rss feed correctly",
		"./testdata/rss.xml",
		map[string]RssItem{
			"sub1-title": {
				Title:       "sub1-title",
				Source:      "main-title",
				SourceURL:   "main-link",
				Description: "sub1-desc",
				Link:        "sub1-link",
			},
			"sub2-title": {
				Title:       "sub2-title",
				Source:      "main-title",
				SourceURL:   "main-link",
				Description: "sub2-desc",
				Link:        "sub2-link",
			},
		},
		2,
	},
}

func TestProcessUrl(t *testing.T) {
	for _, tt := range testProcessUrl {
		t.Run(tt.name, func(t *testing.T) {
			server := startMockServer(t, tt.xmlInputFilePath)
			defer server.Close()

			u := make(chan string, 1)
			r := make(chan RssItem, 2)

			go func() {
				defer func() {
					fmt.Println("hello")
					close(r)
				}()
				processUrl(r, u)
			}()

			u <- server.URL
			close(u)

			var items []RssItem
			for item := range r {
				items = append(items, item)
			}

			if tt.expectedItemsSize != len(items) {
				t.Errorf("Size of expected item %d does not match actual %d", tt.expectedItemsSize, len(items))
				return
			}
			validateItems(t, items, tt.expectedRssItems)
		})
	}
}

var testParse = []struct {
	name              string
	xmlInputFilePath  string
	expectedRssItems  map[string]RssItem
	expectedItemsSize int
}{
	{
		"Test should parse rss feed items correctly",
		"./testdata/rss.xml",
		map[string]RssItem{
			"sub1-title": {
				Title:       "sub1-title",
				Source:      "main-title",
				SourceURL:   "main-link",
				Description: "sub1-desc",
				Link:        "sub1-link",
			},
			"sub2-title": {
				Title:       "sub2-title",
				Source:      "main-title",
				SourceURL:   "main-link",
				Description: "sub2-desc",
				Link:        "sub2-link",
			},
		},
		2,
	},
}

func TestParse(t *testing.T) {
	for _, tt := range testParse {
		t.Run(tt.name, func(t *testing.T) {
			server := startMockServer(t, tt.xmlInputFilePath)
			defer server.Close()

			items := Parse([]string{server.URL})

			if tt.expectedItemsSize != len(items) {
				t.Errorf("Size of expected item %d does not match actual %d", tt.expectedItemsSize, len(items))
				return
			}
			validateItems(t, items, tt.expectedRssItems)
		})
	}
}

func TestParseWithEmptyUrls(t *testing.T) {
	items := Parse([]string{})

	if len(items) != 0 {
		t.Errorf("Size of expected items 0 does not match actual %d", len(items))
		return
	}
}

// could be extended to return error response but the returning errors should be introduced in rss package
func startMockServer(t *testing.T, responseFilePath string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := os.ReadFile(responseFilePath)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := w.Write(b); err != nil {
			t.Fatal(err)
		}
	}))
	return server
}

func validateItems(t *testing.T, items []RssItem, expectedRssItems map[string]RssItem) {
	for _, item := range items {
		v, ok := expectedRssItems[item.Title]
		if !ok {
			t.Errorf("Item %s is missing from expected items", item.Title)
			return
		}
		if !reflect.DeepEqual(expectedRssItems[item.Title], item) {
			t.Errorf("Item %s is not the same as the expected one", v.Title)
			return
		}
	}
}
