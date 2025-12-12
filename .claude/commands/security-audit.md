Perform a comprehensive security audit of the go-logger library.

Invoke the @security-auditor agent to conduct a thorough security review focusing on:

1. **Application Code Security**
   - Review logger.go for security vulnerabilities
   - Check for path traversal issues in file handling
   - Analyze log injection vulnerabilities
   - Verify secure defaults in configuration

2. **Dependencies & Vendor Security**
   - Run `govulncheck ./...` to check for known vulnerabilities
   - Analyze zerolog and lumberjack dependencies
   - Check go.mod/go.sum for outdated packages
   - Review transitive dependencies for CVEs

3. **Configuration Security**
   - Review default configuration security
   - Check file permission handling
   - Analyze log rotation security implications
   - Verify no sensitive data logging patterns

4. **Test Security**
   - Review test code for security anti-patterns
   - Check for hardcoded credentials or secrets
   - Verify proper cleanup of test artifacts

The audit should generate detailed reports in the `.audit/` directory:
- `security-audit-application-code.md` - Source code vulnerabilities
- `security-audit-dependencies.md` - Dependency CVEs and risks
- `security-audit-configuration.md` - Configuration security issues
- `security-audit-summary.md` - Executive summary

All reports should be saved to `.audit/` (which is gitignored).

After the audit, provide:
1. Summary of findings by severity (Critical/High/Medium/Low)
2. Prioritized remediation roadmap
3. Offer to help fix critical and high severity issues
