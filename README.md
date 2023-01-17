# RSS Reader

## Table Of Content
1. [Package `rss`](#rss-package)
2. [Package `meta`](#meta-package)
3. [Usage](#usage)
4. [Developer Improvements Notes](#developer-improvement-notes)
5. [Testing](#testing)

## `rss` Package
This package provides a way to parse asynchronous multiple RSS feeds. It exports the following functions:

```go
Parse(urls []string) []RssItem
```
The function asynchronously processes RSS feeds retrieved using the provided URLs.
They are called using GET request and failed or not correct ones are ignored in the output
The returned `rss.RssItems` are not sorted by published date.

**Input**
| Parameter | Description |
| --- | --- |
| `urls` | An array of strings representing the URLs of the RSS feeds to be parsed. All urls are called with `GET` and should return **XML RSS** | 

**Output**
An array of unordered by date `RssItem` that contain the parsed RSS feed information.

**Example**
```go
package main

import (
	"fmt"
	"rss"
)

func main() {
	// urls should return correct xml rss feed
	urls := []string{"https://www.example.com/feed1", "https://www.example.com/feed2"}
	items := rss.Parse(urls)

	for _, item := range items {
		fmt.Println(item)
	}
}
```

## `meta` Package
Package `meta` provides a way to parse multiple RSS feeds, writes the parsed data in a stream as JSON and print them on the screen. It exports the following function:

```go
ParseFeedAsJSON(urls []string, w io.Writer, parse RSSParse) error
```
This function parses the RSS feeds specified in the urls parameter, writes the parsed data in a stream as JSON and prints the parsed data on the screen.

**Input**
| Parameter | Description |
| --- | --- |
| `urls` | An array of strings representing the URLs of the RSS feeds to be parsed. All urls are called with `GET` and should return **XML RSS** |
| `w` | Writer stream which parsed JSON data will be passed to |
| `parse` function | A function that takes an array of strings representing the URLs of the RSS feeds to be parsed and returns an array of rss.RssItem structs that contain the parsed RSS feed information. |

**Output**
An error if any occurred, otherwise nil

**Example**
```go
package main

import (
	"fmt"
	"meta"
	"rss"
)

func main() {
   // urls should return correct xml rss feed
	urls := []string{"https://www.example.com/feed1", "https://www.example.com/feed2"}
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
```

## Usage
Start the application by using `go run main.go`

Feed Options:
1. Find real RSS feed API urls
2. You can define function of `type RSSParse func([]string) []rss.RssItem` which could be used as mocking. [Example](main.go)

## Developer Improvement Notes
My solution is partially limited by the task. Some extensions/reseraches which I would have added/done in work environment
 
1. `rss.Parse` package: 
   - function should return error or implement alert mechanism using Sentry (Alerting/Monitoring) for all failed feed retrievals. Afterward error testing could be added also
   - could accept http client so it could be mockable/configurable
2. I would add go linter as pre-hook (server hook) or in the CI system
3. Go uber leak package could be used in testing for leaks `"go.uber.org/goleak"`
4. Goroutines Analyses - their number depends on various factors/risks/cost/value:
   - Expected load (Load Testing) - (performance should be measured to find the right balance of goroutines/performance/memory usage and utilization)
   - Performance, Memory requirements and limitations
   - Horizontal/Vertical Scaling options should be evaluated
5. Goroutines Number decision - Average API call takes around 200milliseconds. Simulated 5 goroutines with 200 urls which takes around 4s to complete, which gives us 1/4 of the urls.
   - Amdahl's law could be used for optimization also
6. Logs are in JSON format (required by most of the monitoring tools) but logger implementation could vary depending on the context/domain/usecase of the application. I have decided to do singleton logger retrievable through the whole app just for example. Thread safe logger should be used in concurrent application if logs are required. 
7. Testing could be more extensive
8. `rss.Parse` is a function, as I do not find it semantically (OOP) correct to be method


## Testing
I have decided to go with Go native way of testing using Table Driven Testing. Could be done also with BDD (Ginkgo & Gomega)

