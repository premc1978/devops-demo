# Security Policy

## Reporting a Vulnerability
If you discover a security issue, report it privately:
- Email: security@example.com (replace with real contact) or open a private issue for repo maintainers.
- Provide: affected component, reproduction steps, severity, and suggested remediation.

We will acknowledge within 48 hours and provide a timeline for fixes.

## Responsible Disclosure
- Do not publish details publicly until the issue is resolved.
- Rotate or revoke any leaked credentials immediately.

## Secret Handling
- Never commit secrets to source control.
- Use GitHub Secrets or Azure Key Vault for credentials.
- Use automated secret scanners (Gitleaks) in CI — already configured in workflows.

## Recommended Scanners / Checks
- Secrets: Gitleaks
- Static IaC security: Checkov, TFSec
- Code analysis: CodeQL (where applicable)
- Periodic drift detection: terraform plan -detailed-exitcode

## Access & Least Privilege
- Use short-lived credentials where possible.
- Use least-privilege service principals for CI jobs.

## Incident Response
- Triage and assign severity within 24 hours.
- Notify stakeholders and rotate credentials if secret exposure occurred.
- Create a post-incident runbook and record remediation steps.

## Security Contact
Replace with appropriate contact for your org:
- security@example.com
