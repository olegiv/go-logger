# Claude Code Quick Reference for go-logger

This guide shows how to use the Claude Code extensions in this Go library project.

## Agents

Specialized AI agents for specific tasks. Invoke with `@agent-name` in your message.

### @go-test-runner

Expert for running tests, analyzing coverage, and debugging test failures.

**Examples:**
- "**@go-test-runner** run all tests"
- "**@go-test-runner** check test coverage"
- "**@go-test-runner** debug the TestNew failure"
- "**@go-test-runner** add tests for the WithFields method"

**What it does:**
- Runs Go tests with appropriate flags
- Generates and analyzes coverage reports
- Debugs test failures
- Suggests new test cases
- Follows project testing patterns (t.TempDir(), table-driven tests)

---

### @go-code-quality

Expert for code formatting, linting, and quality checks.

**Examples:**
- "**@go-code-quality** run all linters"
- "**@go-code-quality** format the code"
- "**@go-code-quality** fix the vet warnings"
- "**@go-code-quality** check if the code follows Go best practices"

**What it does:**
- Runs gofmt, go vet, golangci-lint
- Auto-fixes formatting issues
- Identifies code quality problems
- Suggests idiomatic Go improvements
- Manages go.mod tidiness

---

### @go-dependency-manager

Expert for managing dependencies, security updates, and versions.

**Examples:**
- "**@go-dependency-manager** update all dependencies"
- "**@go-dependency-manager** check for vulnerable packages"
- "**@go-dependency-manager** upgrade zerolog to latest version"
- "**@go-dependency-manager** what dependencies need updates?"

**What it does:**
- Lists and updates dependencies
- Checks for security vulnerabilities
- Manages version compatibility
- Works with GitHub Actions security scanning
- Tests after dependency updates

---

### @security-auditor

Expert for comprehensive security audits and vulnerability assessments.

**Examples:**
- "**@security-auditor** perform a full security audit"
- "**@security-auditor** check for CVEs in our dependencies"
- "**@security-auditor** review the code for security issues"
- "**@security-auditor** analyze the latest security scan results"

**What it does:**
- Runs govulncheck for Go-specific vulnerabilities
- Scans dependencies for known CVEs
- Reviews code for security anti-patterns
- Generates detailed security reports in `.audit/` directory
- Provides prioritized remediation guidance
- Checks configuration security

---

## Slash Commands

Quick one-line commands for common tasks. Type the command to execute.

### Testing Commands

| Command | What it does |
|---------|--------------|
| **/test** | Run all tests with verbose output |
| **/test-coverage** | Generate coverage report with HTML visualization |
| **/test-race** | Run tests with race detector for concurrency issues |
| **/benchmark** | Run performance benchmarks |

### Code Quality Commands

| Command | What it does |
|---------|--------------|
| **/lint** | Run gofmt, go vet, and golangci-lint |
| **/fmt** | Format all Go files with gofmt |

### Dependency Commands

| Command | What it does |
|---------|--------------|
| **/mod-tidy** | Clean up go.mod and verify dependencies |
| **/deps-update** | Check for and optionally update dependencies |

### Security Commands

| Command | What it does |
|---------|--------------|
| **/security-audit** | Run comprehensive security audit with govulncheck and CVE scanning |

---

## Common Workflows

### Running Tests
```
Option 1: Use slash command
  /test

Option 2: Use agent for detailed analysis
  @go-test-runner run tests and analyze any failures
```

### Checking Code Quality
```
Option 1: Format and lint separately
  /fmt
  /lint

Option 2: Use agent for comprehensive check
  @go-code-quality run all quality checks and fix issues
```

### Updating Dependencies
```
Option 1: Quick tidy
  /mod-tidy

Option 2: Check for updates
  /deps-update

Option 3: Use agent for full management
  @go-dependency-manager check for updates and apply security patches
```

### Test Coverage Analysis
```
Option 1: Quick coverage report
  /test-coverage

Option 2: Detailed analysis with suggestions
  @go-test-runner analyze coverage and suggest new tests for uncovered code
```

### Security Auditing
```
Option 1: Quick security scan
  /security-audit

Option 2: Comprehensive audit with detailed analysis
  @security-auditor perform a full security audit and generate reports

Option 3: Check specific security concerns
  @security-auditor check if our logger handles sensitive data securely
```

---

## Tips

1. **Slash commands** are fast for routine tasks you do frequently
2. **Agents** provide deeper analysis and can handle complex multi-step tasks
3. **Combine them**: Use `/test` to run tests, then ask an agent to analyze failures
4. **Coverage files** (`coverage.out`, `coverage.html`) are gitignored - safe to generate
5. **Race detection** is important for logger libraries (use `/test-race` regularly)
6. **Security reports** are saved to `.audit/` directory (gitignored) - safe to generate locally

---

## Project-Specific Notes

- This is a **library project** (not an application), so there's no build/deploy process
- Tests use **`t.TempDir()`** for isolation - agents understand this pattern
- The project has **GitHub Actions** for CodeQL and dependency review
- Coverage reports generate **HTML output** you can open in a browser
- Main dependencies are **zerolog** and **lumberjack** - agents know these

---

## Getting Help

- See [CLAUDE.md](../CLAUDE.md) for detailed project architecture
- See [README.md](../README.md) for library usage examples
- Ask agents directly: "What can you help with?" or "What commands are available?"
