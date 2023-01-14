# RSS Reader

## Developer Extension Notes
My solution is limited by the task. Extension which I would have added 
1. `rss.Parse` function should return error or implement alert mechanism using Sentry (Alerting/Monitoring) for all failed feed retrievals.
2. I would add go linter as pre-hook (server hook) or in the CI system
3. Go uber leak package could be used in testing for leaks `"go.uber.org/goleak"`