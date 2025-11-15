# go-logger

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
| `MaxSizeMB` | int | `10` | Maximum size of a log file in megabytes before rotation |
| `MaxBackups` | int | `5` | Maximum number of old log files to retain |
| `Console` | bool | `false` | Enable console output in addition to file logging |

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
    MaxSizeMB:  100,
    MaxBackups: 10,
    Console:    false, // File only in production
})
```

### Development Configuration

```go
log := logger.New(logger.Config{
    Level:   "debug",
    Console: true, // See logs in terminal during development
})
```

## Log Output Format

### File Output
Logs are written to `{LogDir}/go.log` in JSON format:
```json
{"level":"info","time":"2025-11-15T10:30:00Z","caller":"main.go:42","message":"Application started"}
```

### Console Output
When `Console` is enabled, logs are displayed in a human-readable format:
```
2025-11-15 10:30:00 INF Application started caller=main.go:42
```

## Testing

Run the test suite:

```bash
go test -v
```

Run tests with coverage:

```bash
go test -cover
```

## Dependencies

- [zerolog](https://github.com/rs/zerolog) - Zero allocation JSON logger
- [lumberjack](https://github.com/natefinch/lumberjack) - Log file rotation

## License

MIT License - Copyright (c) 2025 Oleg Ivanchenko

See [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
