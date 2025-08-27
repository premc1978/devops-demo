# Deployment Guide

This document explains how deployments are performed for each environment and outlines promotion, approval, and rollback procedures.

## Environments
- `qa` — automated validation and integration tests (PRs to `qa`)
- `dev` — integration/staging; manual review required (PRs to `dev`)
- `main` (`prod`) — production with strict approvals (PRs from `dev` only)

## Promotion Flow (automated)
1. Create PR -> `qa` (runs secrets-scan, checkov, tfsec, terraform init/plan, terratest dry-run).
2. On successful QA checks the workflow auto-approves and merges to `qa`, then creates PR -> `dev`.
3. PR -> `dev` runs full checks; reviewers must approve before merge.
4. After approved merge to `dev`, create PR -> `main` and require prod approvals via GitHub Environments.

## Terraform execution policy
- On PRs: run `terraform init`, `terraform validate`, `terraform plan` only.
- On branch pushes/merges to environment branches: run `terraform apply` (workflows gate apply to push events).
- Use `-var-file` per environment and remote state backend with locking.

Example command (manual):
```
cd tf/environments/dev
terraform init -var-file="terraform.tfvars"
terraform plan -var-file="terraform.tfvars"
# after approvals and merge to branch:
terraform apply -auto-approve -var-file="terraform.tfvars"
```

## Workflows & CI
- Workflows set `TF_TEST_DIR` and `TF_VAR_FILE` for Terratest and Terraform.
- Terratest runs a dry-run plan by default to avoid destructive CI actions.

## Rollback strategy
1. If deployment causes issues, revert the merge commit and push revert PR to the target branch.
2. Alternatively, use Terraform to change resources to known-good config (apply a previous tag/commit).
3. Document and run incident-specific rollback steps in runbook.

## Runbook checklist (pre-deploy)
- Ensure remote state is healthy and unlocked.
- Verify service principal used by CI has correct permissions and no excessive privileges.
- Confirm backups and monitoring/alerts are active.

## Post-deploy validation
- Run smoke tests and health checks.
- Confirm metrics and logs for new resources.
- Notify stakeholders and update changelog.

## Notes
- For destructive integration tests, use ephemeral environments with strict cost and approval controls.
