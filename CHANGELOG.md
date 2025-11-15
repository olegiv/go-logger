# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.0] - 2025-11-15

### Added

- Initial release of go-logger
- Core `Logger` struct wrapping zerolog with enhanced functionality
- `Config` struct for flexible logger configuration
- `New()` factory function with multi-writer support (file + console)
- Automatic log rotation using lumberjack
  - Configurable max file size (default: 10MB)
  - Configurable max backups (default: 5)
  - Automatic cleanup of logs older than 30 days
- Structured logging with zerolog integration
- Context methods for immutable logger instances:
  - `WithField()` - Add single field to log context
  - `WithFields()` - Add multiple fields to log context
  - `WithError()` - Add error to log context
- Support for standard log levels: debug, info, warn, error, fatal, panic
- Case-insensitive log level parsing
- Graceful fallback to stderr if log directory creation fails
- Console output support with human-readable timestamps
- Automatic caller information in all logs
- Comprehensive test suite with >80% coverage
- MIT License
- Complete documentation and README
- GitHub Actions workflows for CodeQL and Dependency Review

[0.0.0]: https://github.com/olegiv/go-logger/releases/tag/v0.0.0
