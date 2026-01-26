# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.1] - 2026-01-26

### Changed

- Updated Go version from 1.25.5 to 1.25.6
- Updated golang.org/x/sys from v0.39.0 to v0.40.0
- Updated copyright year range to 2025-2026
- Migrated Claude Code shared tools to git submodule

## [0.2.0] - 2025-12-12

### Security

- Fixed path traversal vulnerability in LogDir configuration (SEC-001)
- Fixed insecure default directory permissions, changed from 0755 to 0750 (SEC-002)
- Fixed global state pollution via SetGlobalLevel (SEC-003)
- Fixed silent fallback to stderr with structured error logging (SEC-004)
- Added Dependabot for automated dependency monitoring (SEC-005)
- Added optional caller info control via `DisableCaller` config (SEC-006)
- Implemented proper Close() method with file writer cleanup (SEC-007)
- Updated Go toolchain from 1.25.4 to 1.25.5 (SEC-008)

### Added

- `DirMode` config field for custom directory permissions (default: 0750)
- `DisableCaller` config field to disable file path disclosure
- Dependabot configuration for weekly dependency checks
- Dependency monitoring workflow with govulncheck
- 15+ new security-focused tests

### Changed

- Per-logger level configuration instead of global state
- Close() method now properly flushes and closes log files

## [0.1.0] - 2025-11-16

### Added

- Configurable log filename via `Filename` field in `Config` struct
- Default filename remains "go.log" for backward compatibility
- New tests: `TestFilenameDefault` and `TestCustomFilename`
- pkg.go.dev reference badge to README

### Changed

- Updated documentation to reflect configurable filename option

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

[0.2.1]: https://github.com/olegiv/go-logger/releases/tag/v0.2.1
[0.2.0]: https://github.com/olegiv/go-logger/releases/tag/v0.2.0
[0.1.0]: https://github.com/olegiv/go-logger/releases/tag/v0.1.0
[0.0.0]: https://github.com/olegiv/go-logger/releases/tag/v0.0.0
