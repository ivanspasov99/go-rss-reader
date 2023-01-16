package rss

import (
	"encoding/xml"
	"net/http"
	"sync/atomic"
	"time"
)

type RssItem struct {
	Title       string `xml:"title"`
	Source      string
	SourceURL   string
	Link        string `xml:"link"`
	PublishDate time.Time
	Description string `xml:"description"`
}

type channel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	ItemList    []RssItem `xml:"item"`
}

type rss struct {
	Channel channel `xml:"channel"`
}

// Parse asynchronously process rss feeds retrieved using urls
// Returned RssItems are not sorted by Date
func Parse(urls []string) []RssItem {
	var items []RssItem

	u := make(chan string, len(urls))
	r := make(chan RssItem)

	// the following number of goroutines with average request of 200 millisecond would parse all urls for ~5s
	nWorkers := calculateWorkers(len(urls))
	workers := int64(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func() {
			defer func() {
				// Last one out closes channel
				if atomic.AddInt64(&workers, -1) == 0 {
					close(r)
				}
			}()
			processUrl(r, u)
		}()
	}

	for _, v := range urls {
		u <- v
	}
	close(u)

	for item := range r {
		items = append(items, item)
	}
	return items
}

func processUrl(rssItems chan<- RssItem, urls <-chan string) {
	for url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			// Send to Sentry (Monitoring/Alerting tool) for example if no error option appear
			// use thread safe logger
			continue
		}

		if resp.StatusCode != http.StatusOK {
			continue
		}

		var rss rss
		if err := xml.NewDecoder(resp.Body).Decode(&rss); err != nil {
			resp.Body.Close()
			continue
		}

		for _, item := range rss.Channel.ItemList {
			item.Source = rss.Channel.Title
			item.SourceURL = rss.Channel.Link
			rssItems <- item
		}
		resp.Body.Close()
	}
}

func calculateWorkers(urls int) int {
	if urls < 4 {
		return 1
	}
	return urls / 4
}
