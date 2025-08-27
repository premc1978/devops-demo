# Changelog

All notable changes to this repository should be documented in this file.

## Format & Guidance
- Follow [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) style.
- Use semantic versioning (if you publish releases) or date-based headings for infra changes.

### Unreleased
- Describe pending changes and PR references.

### Example entry
#### [Unreleased]
- Added TF variable for resource tagging (#123)
- Updated TFSec policies and fixed high severity findings (#130)

#### 2025-08-01 - v1.2.0
- Promote infrastructure change: adjusted VM sizes for prod (PR #120)
- Updated modules/networking to use new subnet design (PR #118)

## How to update
- Add entries under `Unreleased` when a PR merges.
- On release, move `Unreleased` to a dated heading or version and add release notes.

## Maintainers
- Keep changelog concise: write user-facing impact and references to PRs/issues.
