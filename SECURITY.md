# Security Policy

## Supported Versions

| Version | Supported          | Security Updates |
| ------- | ------------------ | ---------------- |
| 0.1.x   | :white_check_mark: | Yes              |
| < 0.1   | :x:                | No               |

## Reporting a Vulnerability

If you discover a security vulnerability in go-logger, please report it by emailing:

**security@[your-domain].com** or create a private security advisory on GitHub.

Please do NOT open a public issue for security vulnerabilities.

### What to Include

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if available)

### Response Timeline

- **Initial Response**: Within 48 hours
- **Status Update**: Within 7 days
- **Fix Timeline**: Based on severity
  - Critical: 24-48 hours
  - High: 1 week
  - Medium: 2 weeks
  - Low: Next release

## Security Best Practices

### Secure Configuration

#### Production Environment

```go
import (
    "github.com/olegiv/go-logger"
)

cfg := logger.Config{
    Level:      "warn",           // Only warnings and errors
    LogDir:     "/var/log/myapp", // Dedicated directory
    Filename:   "app.log",
    MaxSizeMB:  50,
    MaxBackups: 10,
    Console:    false,            // No console output
    DirMode:    0750,             // Restrictive permissions
}

log := logger.New(cfg)
```

#### Development Environment

```go
cfg := logger.Config{
    Level:      "debug",
    LogDir:     "./logs",
    Console:    true,
    DirMode:    0750,
}

log := logger.New(cfg)
```

### Path Security

Always use absolute paths or validate relative paths:

```go
// GOOD: Absolute path
cfg := logger.Config{
    LogDir: "/var/log/myapp",
}

// GOOD: Validated relative path
cfg := logger.Config{
    LogDir: "./logs",
}

// BAD: User-controlled path without validation
cfg := logger.Config{
    LogDir: userInput,  // NEVER DO THIS
}
```

### File Permissions

Use restrictive permissions for sensitive logs:

```go
// Maximum security (owner only)
cfg := logger.Config{
    DirMode: 0700,  // rwx------
}

// Balanced (owner + group)
cfg := logger.Config{
    DirMode: 0750,  // rwxr-x---
}

// Less secure (world-readable)
cfg := logger.Config{
    DirMode: 0755,  // rwxr-xr-x (not recommended)
}
```

### What NOT to Log

Never log sensitive information:

```go
// BAD - Don't do this!
log.Info().
    Str("password", user.Password).           // NEVER log passwords
    Str("api_key", apiKey).                   // NEVER log API keys
    Str("credit_card", card.Number).          // NEVER log PII
    Str("ssn", user.SSN).                     // NEVER log SSN
    Msg("User login")

// GOOD - Log safely
log.Info().
    Str("username", user.Username).           // OK: username
    Str("user_id", user.ID).                  // OK: user ID
    Str("ip_address", hashIP(r.RemoteAddr)).  // OK: hashed IP
    Msg("User login")
```

### Log Sanitization

Always sanitize user input before logging:

```go
import "strings"

func sanitizeForLog(input string) string {
    // Remove control characters
    input = strings.Map(func(r rune) rune {
        if r < 32 || r == 127 {
            return -1
        }
        return r
    }, input)

    // Truncate long inputs
    if len(input) > 200 {
        input = input[:200] + "...[truncated]"
    }

    return input
}

log.Info().
    Str("user_input", sanitizeForLog(userInput)).
    Msg("Processing request")
```

## Known Vulnerabilities

### Current Issues

None reported as of 2025-12-09.

### Fixed Issues

- None (first release)

## Security Audit

Last security audit: 2025-12-09

Audit findings available in:
- `security-audit-summary.md`
- `security-audit-application-code.md`
- `security-audit-dependencies.md`
- `security-audit-configuration.md`

## Dependency Security

### Direct Dependencies

- **zerolog** v1.34.0 - No known vulnerabilities
- **lumberjack.v2** v2.2.1 - No known vulnerabilities (unmaintained)

### Monitoring

We use the following tools to monitor dependencies:
- govulncheck (Go official vulnerability scanner)
- GitHub Dependabot
- GitHub Security Advisories

## Compliance

### Logging Standards

This library supports compliance with:
- **PCI-DSS**: Requirement 10 (Logging and Monitoring)
- **HIPAA**: Security Rule (Audit Controls)
- **GDPR**: Article 32 (Security of Processing)
- **SOC 2**: Logging and Monitoring Controls

### Recommendations

1. **Encrypt logs at rest** - Use filesystem encryption
2. **Restrict access** - Use DirMode 0700 or 0750
3. **Rotate logs** - Configure MaxBackups and MaxAge
4. **Monitor access** - Use file integrity monitoring
5. **Backup logs** - Implement log backup procedures

## Security Checklist

When using go-logger in production:

- [ ] Use restrictive directory permissions (0750 or 0700)
- [ ] Never log passwords, API keys, or PII
- [ ] Validate all configuration inputs
- [ ] Use absolute paths or validated relative paths
- [ ] Enable log compression for space efficiency
- [ ] Configure appropriate MaxAge for retention
- [ ] Implement log backup procedures
- [ ] Monitor for unauthorized access
- [ ] Regularly update dependencies
- [ ] Review logs for sensitive data leaks
- [ ] Implement log encryption if required
- [ ] Set up alerts for security events

## Contact

For security concerns, contact:
- Email: security@[your-domain].com
- GitHub: https://github.com/olegiv/go-logger/security/advisories

## Acknowledgments

We appreciate security researchers who responsibly disclose vulnerabilities.

Hall of Fame:
- (None yet)

## Updates

This security policy is reviewed quarterly and updated as needed.

Last updated: 2025-12-09
