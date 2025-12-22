# go-logger

[![CodeQL Advanced](https://github.com/olegiv/go-logger/actions/workflows/codeql.yml/badge.svg)](https://github.com/olegiv/go-logger/actions/workflows/codeql.yml)
[![Dependency review](https://github.com/olegiv/go-logger/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/olegiv/go-logger/actions/workflows/dependency-review.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/olegiv/go-logger.svg)](https://pkg.go.dev/github.com/olegiv/go-logger)

A lightweight, production-ready Go logging library built on top of [zerolog](https://github.com/rs/zerolog) with automatic log rotation using [lumberjack](https://github.com/natefinch/lumberjack).

## Features

- **Structured logging** powered by zerolog
- **Automatic log rotation** with configurable size and backup limits
- **Multiple output targets** (file and/or console)
- **Configurable log levels** (debug, info, warn, error)
- **Caller information** automatically included in logs
- **Contextual logging** with field support
- **Timestamp tracking** on all log entries
- **Zero allocation** logging in most cases (thanks to zerolog)
- **Security hardened** with path traversal protection and secure directory permissions

## Installation

```bash
go get github.com/olegiv/go-logger
```

## Quick Start

```go
package main

import (
    "github.com/olegiv/go-logger"
)

func main() {
    // Create logger with default configuration
    log := logger.New(logger.Config{
        Level:      "info",
        LogDir:     "./logs",
        Filename:   "app.log",
        MaxSizeMB:  10,
        MaxBackups: 5,
        Console:    true,
    })
    defer log.Close()

    // Simple logging
    log.Info().Msg("Application started")
    log.Debug().Msg("This won't show at info level")
    log.Warn().Msg("Warning message")
    log.Error().Msg("Error message")

    // Structured logging
    log.Info().
        Str("user", "alice").
        Int("age", 30).
        Msg("User logged in")
}
```

## Configuration

The `Config` struct supports the following options:

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `Level` | string | `"info"` | Log level: `debug`, `info`, `warn`, or `error` |
| `LogDir` | string | `"./logs"` | Directory where log files are stored |
| `Filename` | string | `"go.log"` | Name of the log file |
| `MaxSizeMB` | int | `10` | Maximum size of a log file in megabytes before rotation |
| `MaxBackups` | int | `5` | Maximum number of old log files to retain |
| `Console` | bool | `false` | Enable console output in addition to file logging |
| `DirMode` | os.FileMode | `0750` | Directory permissions (rwxr-x---) for log directory |
| `DisableCaller` | bool | `false` | Disable caller info (file:line) in logs for enhanced privacy |

### Log Rotation

Logs are automatically rotated when they reach the `MaxSizeMB` size limit. Old logs are compressed and retained according to the `MaxBackups` setting. Logs older than 30 days are automatically deleted.

## Usage Examples

### Basic Logging

```go
log := logger.New(logger.Config{
    Level:   "debug",
    Console: true,
})

log.Debug().Msg("Debug message")
log.Info().Msg("Info message")
log.Warn().Msg("Warning message")
log.Error().Msg("Error message")
```

### Contextual Logging with Fields

```go
// Add a single field
logWithUser := log.WithField("user_id", "12345")
logWithUser.Info().Msg("User action")

// Add multiple fields
logWithContext := log.WithFields(map[string]interface{}{
    "request_id": "abc-123",
    "ip":         "192.168.1.1",
    "method":     "POST",
})
logWithContext.Info().Msg("Request received")
```

### Error Logging

```go
err := someFunction()
if err != nil {
    log.WithError(err).Error().Msg("Operation failed")
}
```

### Production Configuration

```go
log := logger.New(logger.Config{
    Level:      "info",
    LogDir:     "/var/log/myapp",
    Filename:   "myapp.log",
    MaxSizeMB:  100,
    MaxBackups: 10,
    Console:    false, // File only in production
    DirMode:    0700,  // Owner-only access (rwx------) for sensitive logs
})
```

### Development Configuration

```go
log := logger.New(logger.Config{
    Level:   "debug",
    Console: true, // See logs in terminal during development
})
```

## Security

This library implements security best practices to protect against common vulnerabilities:

### Path Traversal Protection

All file paths are automatically sanitized to prevent directory traversal attacks:

- **Path Cleaning**: Paths are normalized using `filepath.Clean()` to remove `.` and `..` segments
- **Traversal Detection**: Attempts to use `..` in paths are detected and blocked
- **Filename Validation**: Filenames cannot contain path separators (`/` or `\`)
- **Safe Fallback**: Invalid paths trigger automatic fallback to stderr with security warnings

```go
// ✅ Valid paths - these work correctly
logger.New(logger.Config{LogDir: "./logs"})
logger.New(logger.Config{LogDir: "./logs/app/debug"})
logger.New(logger.Config{Filename: "app.v2.log"})

// ❌ Blocked paths - these fall back to stderr
logger.New(logger.Config{LogDir: "../../etc"})           // Path traversal
logger.New(logger.Config{Filename: "../etc/passwd"})     // Path in filename
```

### Secure Directory Permissions

Log directories are created with restrictive permissions by default:

- **Default**: `0750` (rwxr-x---) - Owner and group only
- **Configurable**: Set custom permissions via `DirMode`
- **Recommended for sensitive logs**: `0700` (rwx------) - Owner only

```go
// Maximum security - owner-only access
logger.New(logger.Config{
    LogDir:  "/var/log/myapp",
    DirMode: 0700,
})
```

### Privacy: Caller Information Control

By default, the logger includes caller information (source file and line number) in all log entries for debugging purposes:

```json
{"level":"info","time":"2025-11-15T10:30:00Z","caller":"main.go:42","message":"Application started"}
```

For enhanced privacy and security, you can disable caller information to prevent exposing internal application structure:

```go
// Disable caller info for production logs
logger.New(logger.Config{
    LogDir:        "/var/log/myapp",
    DisableCaller: true,  // Omit file paths from logs
})
```

**When to disable caller info**:
- Production environments where logs may be accessed by untrusted parties
- When log files could be exposed (e.g., public S3 buckets, log aggregation services)
- Compliance requirements (GDPR, HIPAA) that prohibit exposing internal paths
- To prevent reconnaissance attacks that use file structure information

**When to keep caller info enabled** (default):
- Development and debugging
- Controlled production environments with secure log access
- When troubleshooting issues requires stack trace context

### Security Warnings

When security violations are detected, warnings are logged to stderr:

```json
{
  "level": "error",
  "security_warning": "path traversal detected in LogDir: ../../etc",
  "message": "SECURITY: Invalid logger configuration, falling back to stderr"
}
```

## Log Output Format

### File Output
Logs are written to `{LogDir}/{Filename}` in JSON format:
```json
{"level":"info","time":"2025-11-15T10:30:00Z","caller":"main.go:42","message":"Application started"}
```

### Console Output
When `Console` is enabled, logs are displayed in a human-readable format:
```
2025-11-15 10:30:00 INF Application started caller=main.go:42
```

## Development with Claude Code

This project includes Claude Code extensions for enhanced development workflow. Some tools are shared via a [git submodule](https://github.com/olegiv/claude-code-support-tools).

### Cloning for Development

```bash
# Clone with submodules
git clone --recurse-submodules https://github.com/olegiv/go-logger.git

# Or initialize after clone
git submodule update --init --recursive

# Update submodule to latest
git submodule update --remote .claude/shared
```

> **Note**: Some files in `.claude/agents/` and `.claude/commands/` are symlinks to the shared submodule. These may appear as broken links in GitHub's web UI, but work correctly when cloned locally.

### Quick Commands

```bash
/test              # Run all tests
/test-coverage     # Generate coverage report
/test-race         # Run with race detector
/lint              # Check code quality
/fmt               # Format all files
/mod-tidy          # Clean up dependencies
/deps-update       # Update dependencies
/benchmark         # Run benchmarks
/security-audit    # Run security audit and CVE scan
/commit-prepare    # Review changes and prepare commit message
/commit-do         # Create commit with prepared message
/setup-project-tools  # Generate project-specific Claude Code tools
```

### Specialized Agents

- **@go-test-runner** - Testing and coverage analysis
- **@go-code-quality** - Linting and formatting
- **@go-dependency-manager** - Dependency management
- **@security-auditor** - Security audits and vulnerability scanning
- **@project-architect** - Generate tailored Claude Code agents/commands
- **@code-quality-auditor** - Scan and fix code quality issues

See [CLAUDE.md](CLAUDE.md) for detailed usage.

## Testing

Run the test suite:

```bash
go test -v
```

Run tests with coverage:

```bash
go test -cover
```

Generate HTML coverage report:

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

## Dependencies

### Direct Dependencies

- **[zerolog](https://github.com/rs/zerolog)** v1.34.0 - Zero allocation JSON logger
  - Status: ✅ Actively maintained
  - Last updated: March 2025

- **[lumberjack](https://github.com/natefinch/lumberjack)** v2.2.1 - Log file rotation
  - Status: ⚠️ Unmaintained (last release: Feb 2023)
  - Security: No known CVEs
  - Monitoring: Automated weekly checks via Dependabot and GitHub Actions

### Dependency Status

The lumberjack.v2 library is currently **unmaintained** but remains **stable and secure** with no known vulnerabilities. We have implemented automated monitoring to track its status:

- **Dependabot**: Weekly checks for new versions
- **Dependency Monitor**: Weekly workflow checking repository activity
- **govulncheck**: Weekly vulnerability scans

**Long-term plan**: We are monitoring the situation and will evaluate actively maintained alternatives (such as [log/slog](https://pkg.go.dev/log/slog) from Go 1.21+) if the dependency becomes problematic or security issues arise.

**Current recommendation**: Safe to use. The library is feature-complete, stable, and has no known security vulnerabilities.

### Security Notes

- All dependencies are from trusted sources (GitHub, official Go repositories)
- Module checksums are verified via `go.sum`
- CodeQL and Dependency Review workflows run on all pull requests
- Weekly vulnerability scanning with govulncheck

For detailed security audit results, see [.audit/security-audit-dependencies.md](.audit/security-audit-dependencies.md).

## License

MIT License - Copyright (c) 2025 Oleg Ivanchenko

See [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
