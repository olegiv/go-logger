Check for available dependency updates and optionally update them.

Execute:

```bash
cd /Users/olegiv/Desktop/Projects/Go/go-logger

# Show current dependencies
echo "Current dependencies:"
go list -m all

# Check for available updates
echo -e "\nChecking for available updates:"
go list -u -m all
```

Display:
1. Current versions of all dependencies
2. Available updates (if any)
3. Distinction between direct and indirect dependencies

Ask the user if they want to update:
- All dependencies to latest compatible versions
- Specific dependencies only
- Only patch versions (safest)

If user confirms, execute appropriate update command:
```bash
# Update all (minor/patch)
go get -u ./...

# Update patches only
go get -u=patch ./...

# Update specific dependency
go get -u github.com/rs/zerolog
```

After updates, automatically run:
```bash
go mod tidy
go mod verify
go test -v ./...
```

Warn about potential breaking changes for major version updates.
