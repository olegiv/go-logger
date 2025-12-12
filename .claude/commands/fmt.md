Format all Go source files using gofmt.

Execute:

```bash
cd /Users/olegiv/Desktop/Projects/Go/go-logger

# Show which files will be formatted
echo "Files that need formatting:"
gofmt -l .

# Format all files
echo "Formatting all Go files..."
go fmt ./...
```

Report:
- Which files were formatted
- Confirmation that formatting is complete
- Reminder to review changes before committing

Note: This command modifies files in place. The user should review changes using `git diff` before committing.
