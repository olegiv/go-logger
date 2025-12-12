Run Go linters and formatters to check code quality.

Execute the following commands:

```bash
cd /Users/olegiv/Desktop/Projects/Go/go-logger

# Check formatting
echo "Checking code formatting..."
gofmt -l .

# Run go vet
echo "Running go vet..."
go vet ./...

# Run golangci-lint if available
echo "Running golangci-lint..."
if command -v golangci-lint &> /dev/null; then
    golangci-lint run
else
    echo "golangci-lint not installed, skipping"
fi
```

Report:
1. Files that need formatting (if any)
2. Issues found by `go vet`
3. Issues found by `golangci-lint` (if available)

If issues are found, categorize them by:
- Severity (error, warning, info)
- Type (formatting, style, potential bugs, performance)
- Whether they can be auto-fixed

Suggest next steps for fixing issues.
