Run tests with coverage analysis and generate an HTML coverage report.

Execute the following commands:

```bash
cd /Users/olegiv/Desktop/Projects/Go/go-logger
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

Display:
1. Overall coverage percentage
2. Coverage breakdown by function
3. Path to the HTML coverage report (coverage.html)

Analyze the results and:
- Highlight functions with low or no coverage
- Identify untested code paths
- Suggest additional test cases for uncovered areas
- Note that coverage.out and coverage.html are gitignored

After generating the report, inform the user they can open `coverage.html` in a browser to see visual coverage.
