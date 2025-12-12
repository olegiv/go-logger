Run tests with race detection to catch concurrency issues.

Execute:

```bash
cd /Users/olegiv/Desktop/Projects/Go/go-logger
go test -race -v
```

The race detector will identify:
- Data races in concurrent code
- Unsafe concurrent access to shared memory
- Potential race conditions

Note: While this library doesn't have explicit concurrency features, race detection ensures the logger is safe for concurrent use (which is a common use case for loggers).

Report any race conditions found and suggest fixes.
