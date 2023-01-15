package rss

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

type RssItem struct {
	Title     string `xml:"title"`
	Source    string
	SourceURL string
	Link      string `xml:"link"`
	// TODO handle time
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

// Parse asynchronously process rss feeds retrieved using input urls
// Returned RssItems are not sorted by Date
func Parse(urls []string) []RssItem {
	var items []RssItem

	u := make(chan string, len(urls))
	r := make(chan RssItem)

	// the following number of goroutines with average request of 200 millisecond would parse all urls for ~5s
	nWorkers := len(urls) / 4
	workers := int32(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func() {
			defer func() {
				// Last one out closes does -1 to workers for every finished go routine
				if atomic.AddInt32(&workers, -1) == 0 {
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
		fmt.Println("url processed", url)
		resp, err := http.Get(url)
		if err != nil {
			// Log or send to Sentry (Monitoring/Alerting tool) for example if no error option appear
			// return error if possible
			fmt.Println("Get request have failed ", err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Status code is not OK", resp.StatusCode)
			return
		}

		var rss rss
		if err := xml.NewDecoder(resp.Body).Decode(&rss); err != nil {
			resp.Body.Close()
			fmt.Println("Decode xml has failed", err)
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
