# DevOps Demo - Infrastructure as Code CI/CD Pipeline

This repository demonstrates enterprise-grade CI/CD patterns for Infrastructure as Code (IaC) using Terraform, GitHub Actions, and industry best practices for scalable, secure, and maintainable infrastructure deployments.

## Table of Contents

1. [Pipeline Architecture](#pipeline-architecture)
2. [Pipeline Patterns That Scale](#pipeline-patterns-that-scale)
3. [Best-Practice CI/CD Structure](#best-practice-cicd-structure)
4. [Approval Workflow Design](#approval-workflow-design)
5. [Environment Promotion Strategy](#environment-promotion-strategy)
6. [Security & Compliance](#security--compliance)
7. [Directory Structure](#directory-structure)
8. [Getting Started](#getting-started)
9. [Common Patterns & Anti-Patterns](#common-patterns--anti-patterns)

## Pipeline Architecture

### Environment Flow
```
Feature Branch → QA → Dev → Prod
```

- **QA Environment**: Automated testing and validation
- **Dev Environment**: Integration testing and staging
- **Prod Environment**: Production deployments with manual approval

### Branch Strategy
- `main` branch → Production environment
- `dev` branch → Development environment  
- `qa` branch → QA environment
- Feature branches → PR-based testing

## Pipeline Patterns That Scale

### 1. **Modular Job Design**
```yaml
# Each job has a single responsibility
jobs:
  secrets-scan:    # Security scanning
  static-analysis: # Code quality (Checkov, TFSec)
  terraform:       # Infrastructure validation & deployment
  integration:     # End-to-end testing (Terratest)
```

### 2. **Environment-Specific Workflows**
- **Separate workflow files** per environment prevent cross-environment contamination
- **Consistent job structure** across environments ensures reliability
- **Environment-specific configurations** through directory structure

### 3. **Fail-Fast Strategy**
```yaml
# Dependencies ensure early failure detection
terraform:
  needs: [secrets-scan, checkov, tfsec]
terratest:
  needs: terraform
```

### 4. **Parallel Execution**
Static analysis jobs run in parallel for faster feedback:
- Secrets scanning
- Security compliance (Checkov)  
- Terraform security (TFSec)

## Best-Practice CI/CD Structure

### Standard Pipeline Flow
```
1. Code Commit/PR
2. Secrets Scan (Gitleaks)
3. Static Analysis (Checkov, TFSec)
4. Terraform Validate & Plan  
5. Integration Tests (Terratest)
6. Approval Gate (Environment-specific)
7. Terraform Apply
8. Post-deployment Validation
```

### Speed Optimization Strategies

#### 1. **Caching Strategy**
```yaml
# Cache Terraform providers and modules
- uses: actions/cache@v3
  with:
    path: |
      ~/.terraform.d/
      .terraform/
    key: terraform-${{ hashFiles('**/.terraform.lock.hcl') }}
```

#### 2. **Selective Triggering**
```yaml
# Only run on relevant changes
on:
  push:
    paths: 
      - 'tf/environments/prod/**'
      - '.github/workflows/terraform-prod.yml'
```

#### 3. **Job Optimization**
- **Parallel static analysis** reduces total pipeline time
- **Conditional job execution** based on file changes
- **Lightweight container images** for faster startup

### Modularity Best Practices

#### 1. **Reusable Actions**
```yaml
# Custom composite actions for common tasks
- uses: ./.github/actions/terraform-setup
  with:
    environment: ${{ env.ENVIRONMENT }}
```

#### 2. **Shared Configuration**
```yaml
# Environment-specific variable files
terraform_var_file: "environments/${{ github.ref_name }}.tfvars"
```

## Approval Workflow Design

### Automation vs Control Balance

#### **Auto-Approve Scenarios** ✅
- **QA Environment**: Automated approval after all tests pass
- **Non-critical changes**: Documentation, minor configuration updates
- **Dependency updates**: Automated security patches

#### **Manual Approval Required** 🚫  
- **Production deployments**: Always require human oversight
- **Dev environment**: Team lead approval for architectural changes
- **Security-related changes**: RBAC, networking, encryption modifications

### Implementation Patterns

#### 1. **GitHub Environment Protection**
```yaml
environment:
  name: prod
  # Configured in GitHub repo settings:
  # - Required reviewers
  # - Deployment branches
  # - Wait timer
```

#### 2. **Progressive Approval Gates**
```yaml
# QA: Auto-approve → Dev: Team approval → Prod: Senior approval
approval_levels:
  qa: auto
  dev: team-leads  
  prod: senior-engineers + security-team
```

#### 3. **Conditional Approval Logic**
```yaml
# Smart approval based on change type
if: contains(github.event.pull_request.labels.*.name, 'breaking-change')
# Require additional approvals for breaking changes
```

### Enforcement Mechanisms

#### 1. **Branch Protection Rules**
- Require PR reviews before merging
- Dismiss stale reviews on new commits
- Require status checks to pass

#### 2. **CODEOWNERS Integration**
```
# .github/CODEOWNERS
/tf/environments/prod/     @senior-engineers @security-team
/tf/modules/security/      @security-team
*.tf                       @platform-team
```

## Environment Promotion Strategy

### 1. **PR-Based Promotion**
```
Feature Branch → PR to QA → Auto-merge → PR to Dev → Manual Review → PR to Prod
```

#### Advantages:
- **Audit trail**: Every change tracked through PRs
- **Review gates**: Natural approval points
- **Rollback capability**: Easy revert through Git

### 2. **Automated QA to Dev Promotion**
```yaml
# After QA success, automatically create Dev PR
promote-to-dev:
  needs: [qa-tests-passed]
  steps:
    - name: Create Dev PR
      uses: peter-evans/create-pull-request@v6
```

### 3. **Environment Isolation**
```
tf/environments/
├── qa/         # QA-specific configurations
├── dev/        # Dev-specific configurations  
└── prod/       # Production configurations
```

### 4. **Configuration Management**
```hcl
# environment-specific variables
variable "environment" {
  description = "Environment name"
  type        = string
}

# Environment-specific resource sizing
locals {
  instance_sizes = {
    qa   = "t3.micro"
    dev  = "t3.small" 
    prod = "t3.large"
  }
}
```

### 5. **Drift Detection & Remediation**
```yaml
# Weekly drift detection job
schedule:
  - cron: '0 9 * * MON'  # Every Monday 9 AM
steps:
  - name: Detect Drift
    run: terraform plan -detailed-exitcode
  - name: Alert on Drift  
    if: failure()
    uses: actions/slack-notify
```

## Security & Compliance

### Static Analysis Pipeline
1. **Gitleaks**: Secret detection in code
2. **Checkov**: Policy as Code compliance
3. **TFSec**: Terraform security scanning

### Security Gates
- **No secrets in code**: Automated scanning prevents credential leaks
- **Security policy compliance**: Infrastructure must meet security baselines
- **Least privilege**: IAM roles follow principle of least privilege

## Directory Structure

```
devops-demo/
├── .github/
│   └── workflows/           # Environment-specific workflows
│       ├── terraform-qa.yml
│       ├── terraform-dev.yml
│       ├── terraform-prod.yml
│       └── terraform-other.yml
├── tf/
│   ├── environments/        # Environment-specific configurations
│   │   ├── qa/
│   │   ├── dev/
│   │   └── prod/
│   └── modules/            # Reusable Terraform modules
└── test/                   # Integration tests
    └── terraform_integration_test.go
```

## Getting Started

### Prerequisites
- GitHub repository with Actions enabled
- Azure subscription with service principal
- Terraform >= 1.0
- Go >= 1.19 (for Terratest)

### Setup Steps

1. **Configure GitHub Secrets**
   ```
   AZURE_CREDENTIALS={"clientId":"...","clientSecret":"..."}
   ```

2. **Set up Environment Protection**
   - Navigate to repository Settings → Environments
   - Configure protection rules for each environment
   - Add required reviewers for prod environment

3. **Initialize Terraform Backend**
   ```bash
   # Configure remote state storage
   terraform init -backend-config="key=qa/terraform.tfstate"
   ```

## Common Patterns & Anti-Patterns

### ✅ **Recommended Patterns**

#### 1. **Fail-Fast Pipeline Design**
```yaml
# Run quick, cheap tests first
jobs:
  quick-validation:  # < 30 seconds
  security-scan:     # < 2 minutes  
  terraform-plan:    # < 5 minutes
  integration-tests: # < 15 minutes
```

#### 2. **Environment Parity**
- Consistent resource configurations across environments
- Same pipeline structure for all environments
- Infrastructure-as-Code for all environments

#### 3. **Immutable Infrastructure**
- Replace rather than modify infrastructure
- Version all infrastructure changes
- Use blue-green or rolling deployments

### ❌ **Anti-Patterns to Avoid**

#### 1. **Manual Environment Configuration**
```yaml
# BAD: Manual steps in pipeline
- name: Manually configure load balancer
  run: echo "Please configure the load balancer manually"
```

#### 2. **Shared State Files**
```hcl
# BAD: Single state file for all environments
backend "s3" {
  bucket = "terraform-state"
  key    = "shared/terraform.tfstate"  # DON'T DO THIS
}
```

#### 3. **Overly Complex Approval Chains**
```yaml
# BAD: Too many approval gates slow down delivery
needs: [approval-1, approval-2, approval-3, approval-4]
```

#### 4. **Environment Drift**
```yaml
# BAD: Different pipeline logic per environment
if: env.ENVIRONMENT == 'prod'
  run: special-prod-only-logic
```

### Common Mistakes to Watch For

1. **Secret Management**
   - ❌ Hardcoded secrets in Terraform files
   - ✅ Use GitHub secrets or Azure Key Vault

2. **State Management**  
   - ❌ Local state files
   - ✅ Remote state with locking

3. **Testing Strategy**
   - ❌ Only testing in production
   - ✅ Comprehensive testing in lower environments

4. **Approval Bottlenecks**
   - ❌ Requiring approval for every minor change
   - ✅ Risk-based approval requirements

## Monitoring & Observability

- **Pipeline metrics**: Track deployment frequency, lead time, MTTR
- **Infrastructure monitoring**: CloudWatch, Azure Monitor integration  
- **Cost monitoring**: Track infrastructure costs per environment
- **Security monitoring**: Continuous compliance checking

## Contributing

1. Create feature branch from `main`
2. Make changes in appropriate environment directory
3. Create PR to `qa` branch for testing
4. After QA approval, promote to `dev`
5. Final review and deployment to `prod`

---

*This repository demonstrates production-ready patterns for Infrastructure as Code CI/CD pipelines. Each pattern has been battle-tested in enterprise environments and represents current industry best practices.*

