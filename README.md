# RSS Reader

## Developer Extension Notes
My solution is limited by the task. Extension which I would have added 
1. `rss.Parse` function should return error or implement alert mechanism using Sentry (Alerting/Monitoring) for all failed feed retrievals. Afterward error testing could be added also
2. I would add go linter as pre-hook (server hook) or in the CI system
3. Go uber leak package could be used in testing for leaks `"go.uber.org/goleak"`
4. Goroutines Analyses - their number depends on various factors/risks/cost/value:
   - Expected load (Load Testing) - (performance should be measured to find the right balance of goroutines/performance/memory usage and utilization)
   - Performance, Memory requirements and limitations
   - Horizontal/Vertical Scaling options should be evaluated
5. Goroutines Number RSS reader task decision - Average API call takes around 200milliseconds. Simulated 5 goroutines with 200 urls which takes around 4s to complete, which gives us 1/4 of the url.
>Amdahl's law could be used for optimization also

## Testing
I have decided to go with Go native way of testing using Table Driven Testing. Could be done also with BDD (Ginkgo & Gomega)

