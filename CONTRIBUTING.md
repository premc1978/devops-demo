# Contributing to devops-demo

Thank you for contributing. This document describes the expected contribution workflow, CI/CD checks, testing guidance, and promotion rules for Infrastructure-as-Code changes in this repository.

## Branching & PR flow
- Feature work: create feature branches from `dev` (or `qa` for QA-only experiments).
- Promotion flow:
  - Create PR → `qa` (all automated checks run).
  - After QA checks pass the PR is auto-approved and merged to `qa`, then an automated PR is created from the source branch into `dev`.
  - Create PR → `dev` (full validation runs; PRs to `dev` must be reviewed manually — no auto-approval).
  - After review and merge to `dev`, create PR → `main` (`prod`) only from `dev` (protected — manual approvals required).
- PRs into `main` are restricted: only promote from `dev`. Use the promoted PR flow.

## Required checks (automated)
On PRs the pipeline runs:
1. Secrets scan (Gitleaks)
2. Static scans (Checkov, TFSec)
3. Terraform init/validate/plan
4. Terratest (dry-run plan)
Only after all checks pass will the QA job auto-approve and merge per workflow rules.

## Environments and protections
- Each terraform job binds to a GitHub Environment (`qa`, `dev`, `prod`). Required reviewers and wait timers are enforced there.
- Avoid direct pushes to protected branches (`dev`, `main`).

## Tests & Terratest
- Terratest is configured to run a safe dry-run (terraform init + plan) by default.
- Tests read the target directory from `TF_TEST_DIR` or `TF_ENV`. Workflows set `TF_TEST_DIR` for each environment.
- To run tests locally:
  - Install Go and Terraform.
  - from repo root: `cd test && TF_TEST_DIR=../tf/environments/qa go test -v ./...`
- For destructive integration testing (apply/destroy), use an ephemeral environment and ensure manual approvals and billing considerations.

## Terraform usage
- Keep environment-specific configuration under `tf/environments/<env>/`.
- Use `-var-file` for sensitive or environment variables (workflows set `TF_VAR_FILE`).
- Use remote state with locking; do not commit local state.

## Code style & modules
- Create reusable modules in `tf/modules/`.
- Keep variable and output names consistent and documented in module README.
- Add or update `README.md` for any new module explaining inputs, outputs, and usage.

## PR requirements
- Provide a clear description of the change and impact.
- Include plan output or summary (if applicable).
- Add reviewers and/or ensure CODEOWNERS are set for touched paths.
- Address security and compliance impact (e.g., IAM, network changes).

## Handling secrets & credentials
- Never commit secrets to the repository.
- Use GitHub Secrets or a secrets manager (Azure Key Vault) referenced from workflows.
- If you find a leaked secret, rotate it immediately and open a security issue.

## Reporting vulnerabilities or incidents
- Report infra/security issues privately to the repo owners or security team.
- Mark issues `security` and avoid public disclosure until mitigated.

## Maintainers
- Keep workflows up-to-date and minimal.
- Review `TF_VAR_FILE` usage and remote state configuration when approving infra changes.

Thank you for contributing — following these rules keeps deployments safe, auditable, and repeatable.
