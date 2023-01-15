# RSS Reader

## Usage
- Find real RSS feed API urls which can be used
- Feeds can be retrieved using `rss` package
- You can save feeds and prints to stdout with `meta.ParseFeedAsJSON` using `rss.Parse` function or other custom implemented one 
>For testing purposes you can define function of `type RSSParse func([]string) []rss.RssItem` which could be used as mocking. [Example](pkg/meta/meta_test.go)

## Developer Improvement Notes
My solution is limited by the task. Extension which I would have added 
1. `rss.Parse` package: 
   - function should return error or implement alert mechanism using Sentry (Alerting/Monitoring) for all failed feed retrievals. Afterward error testing could be added also
   - could accept http client so it could be mockable/configurable
2. I would add go linter as pre-hook (server hook) or in the CI system
3. Go uber leak package could be used in testing for leaks `"go.uber.org/goleak"`
4. Goroutines Analyses - their number depends on various factors/risks/cost/value:
   - Expected load (Load Testing) - (performance should be measured to find the right balance of goroutines/performance/memory usage and utilization)
   - Performance, Memory requirements and limitations
   - Horizontal/Vertical Scaling options should be evaluated
5. Goroutines Number RSS reader task decision - Average API call takes around 200milliseconds. Simulated 5 goroutines with 200 urls which takes around 4s to complete, which gives us 1/4 of the url.
   - Amdahl's law could be used for optimization also
6. `meta.SaveToJSON` I prefer using more generalized approach would make the code more extensible and testable:
   - Passing Encoder or Writer Interface
   - Passing enum value and generate object on runtime
7. All logs should be in JSON format (`zerolog` could be used)
8. Testing could be more extensive


## Testing
I have decided to go with Go native way of testing using Table Driven Testing. Could be done also with BDD (Ginkgo & Gomega)

