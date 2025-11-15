# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**go-logger** is a lightweight Go library that wraps [zerolog](https://github.com/rs/zerolog) with automatic log rotation via [lumberjack](https://github.com/natefinch/lumberjack). It's designed as a reusable package, not a standalone application.

## Architecture

### Core Components

- **`logger.go`**: Main library implementation
  - `Logger` struct: Wraps `zerolog.Logger` to provide enhanced functionality
  - `Config` struct: Configuration for log level, directory, rotation settings, and console output
  - `New()`: Factory function that creates a logger with multi-writer support (file + optional console)
  - Context methods: `WithField()`, `WithFields()`, `WithError()` - all return new logger instances preserving immutability

### Key Design Patterns

1. **Wrapper Pattern**: The `Logger` struct embeds `zerolog.Logger`, inheriting all its methods while adding custom functionality
2. **Immutability**: Context methods (`WithField`, `WithFields`, `WithError`) create new logger instances rather than modifying the original
3. **Multi-Writer**: Logs can be written to both file and console simultaneously using `io.MultiWriter`
4. **Graceful Fallback**: If log directory creation fails, the logger falls back to `stderr` instead of failing

### Log Rotation Behavior

- Log files are written to `{LogDir}/go.log`
- Automatic rotation when file reaches `MaxSizeMB` (default: 10MB)
- Keeps `MaxBackups` old files (default: 5)
- Automatically deletes logs older than 30 days (hardcoded in `logger.go:57`)
- Old logs are NOT compressed (`Compress: false` in `logger.go:58`)

## Development Commands

### Testing
```bash
# Run all tests
go test -v

# Run tests with coverage
go test -cover

# Run a specific test
go test -v -run TestNew

# Run tests matching a pattern
go test -v -run TestWith
```

### Building
```bash
# Build the module (verify compilation)
go build

# Download/update dependencies
go mod download

# Tidy dependencies
go mod tidy

# Verify dependencies
go mod verify
```

### Linting and Quality
This project uses GitHub Actions for automated checks:
- **CodeQL**: Runs on push/PR to master, plus weekly schedule
- **Dependency Review**: Runs on PRs to detect vulnerable dependencies

## Testing Philosophy

Tests use `t.TempDir()` to create isolated temporary directories for each test case, ensuring no test pollution. The test suite covers:

- Configuration defaults and validation
- Log level parsing (case-insensitive, handles "warn"/"warning")
- File and directory creation (including nested directories)
- Multi-writer scenarios (file + console)
- Context preservation (WithField/WithFields/WithError return new instances)
- Fallback behavior when directory creation fails

## Configuration Defaults

When creating a logger, these defaults apply:
- `Level`: `"info"` if empty or unrecognized
- `LogDir`: `"./logs"` if empty
- `MaxSizeMB`: `10` if zero
- `MaxBackups`: `5` if zero
- `Console`: `false` by default
- `MaxAge`: 30 days (hardcoded, not configurable)

## Important Implementation Details

1. **Global Level Setting**: `New()` calls `zerolog.SetGlobalLevel()`, which affects ALL zerolog loggers in the process
2. **Caller Information**: Automatically added to all logs via `.Caller()` in `logger.go:80`
3. **Timestamp Format**: File logs use Unix timestamp; console uses `"2006-01-02 15:04:05"` format
4. **No Compression**: Log rotation doesn't compress old files (set `Compress: false`)
5. **Close() is a no-op**: The `Close()` method exists for API completeness but doesn't actually do anything (zerolog doesn't require explicit closing)

## Dependencies

- Go 1.25.4
- `github.com/rs/zerolog` v1.34.0 - Zero-allocation JSON logger
- `gopkg.in/natefinch/lumberjack.v2` v2.2.1 - Log file rotation

## License

MIT License - Copyright (c) 2025 Oleg Ivanchenko
