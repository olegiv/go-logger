# Agent Guidance

## Project

This repository is a reusable Go logging library, not an application.
Use Go 1.27.1 or later; `go.mod` is the source of truth for Go and module versions.
Use idiomatic Go and gofmt. Preserve the exported API unless a change is requested.

- `logger.go` embeds zerolog and adds configuration, file rotation, and context helpers.
- `logger_test.go` covers logging behavior, configuration, fallback, and resource cleanup.
- Context helpers return new loggers, preserve inherited fields and the file writer,
  and must not mutate the original logger.
- Levels are per instance; do not change zerolog's global level.
- Rotation uses lumberjack with 30-day retention and compression disabled.
- Invalid paths or directory failures fall back to stderr. Preserve this behavior.

## Validation

Use Go 1.27.1 and golangci-lint v2.14.0 to match CI. For Go changes, run:

```bash
go build -mod=readonly ./...
go test -mod=readonly -race ./...
go vet -mod=readonly ./...
golangci-lint run ./...
```

For dependency changes, also run `go mod tidy`, `go mod verify`, and
`govulncheck ./...`. For workflow changes, run
`go run github.com/rhysd/actionlint/cmd/actionlint@v1.7.12`.
Always run `git diff --check`. Scope checks to touched components when appropriate.
Test observable behavior; use temporary directories for file tests and close
writers before reading logs. Do not add tests that merely mirror implementation.

GitHub CI includes build, vet, race tests, lint, and vulnerability checks.
CodeQL scans Go and Actions. Dependency review runs on PRs; Dependabot and
dependency monitoring run weekly. Pin workflow Actions to commit SHAs.

## Documentation

Keep README.md, CLAUDE.md, and SECURITY.md dependency references consistent
with go.mod. Describe unreleased changes in CHANGELOG.md without inventing a
release date. Keep documentation consistent with actual logger behavior.

## Git and Reviews

Use `codex/` for new branches. Never overwrite unrelated user changes.
Draft the full commit message and obtain explicit approval before committing.
Use an imperative subject of at most 50 characters, followed by a blank line
and a body wrapped at 72 characters. Do not add AI attribution.
Push only when explicitly authorized; do not merge or publish without authorization.

For PR review findings, use `$pr-fix`: fix the named finding, run one local
`codex exec review --base` before pushing, and allow at most two fix pushes
per PR unless the user says `override`. A P2 gets a one-line fix or a reply,
never a rewrite. Obtain approval before posting replies or resolving threads.
Explicitly requested docs or tests may be updated separately from the finding.

Invoke `$full-branch-audit` only when explicitly named or when asked for a
whole-checkout audit. Audit committed HEAD plus staged, unstaged, and untracked
files through the skill's fresh, blind, ephemeral read-only process. Independently
corroborate candidates and check snapshot stability. Never substitute a diff-only
review or include a whole-repository audit in a PR-findings fix.

## Communication

Assume an experienced developer. Give direct answers in English or Russian,
never French. Ask upfront about material ambiguities; resolve discoverable facts
from the repository first. Prefer current stable software and GUI tools where practical.
Use Europe/Zurich for user-facing dates and times.
