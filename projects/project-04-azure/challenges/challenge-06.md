# Challenge 6: Create a CI/CD pipeline that tests and deploys the application, and introduce infrastructure as code

## Difficulty
Intermediate

## Problem statement and user story
The booking application needs a repeatable deployment flow. Manual deployment is slow, risky, and hard to review. The team also wants infrastructure to be treated as code so deployments can be reproduced and audited.

## Practical outcome
Participants build a CI/CD pipeline and write the infrastructure definitions for deploying their app.

## Learning objectives
- understand the stages of CI/CD
- implement automated validation and deployment
- treat infrastructure as code rather than portal click-throughs
- explain how environment separation improves safety

## Prerequisites and setup
- GitHub or another source control platform
- local environment with a deployment tool or Azure CLI
- a working app and basic test suite

## Azure services under consideration
- GitHub Actions or Azure DevOps
- Azure CLI
- Bicep or Terraform
- Azure Resource Manager patterns

### Verified free-tier / cost notes
- GitHub and Azure DevOps plans vary by account and coverage.
- Azure resource cost still applies when deployment resources are created.
- Check current account-specific pricing and free-tier eligibility before deploying.

## Reference architecture and request/data flow
```text
Code push -> CI validation -> infrastructure apply -> deployment -> smoke tests -> monitoring
```

## Functional and non-functional requirements
Functional:
- code is tested automatically
- app can be deployed to a target environment
- infrastructure definition is version-controlled

Non-functional:
- predictable deployment
- secure handling of secrets
- rollback or redeploy path is documented

## Acceptance criteria
- a commit or pull request triggers automated checks
- deployment works without manual portal changes
- infra is defined as code
- design notes explain environment boundaries

## Minimal implementation path
1. add a basic CI workflow
2. run lint and tests
3. deploy infra using code
4. deploy app to a dev environment
5. run a smoke check

## More advanced implementation path
- separate dev, staging, and production workflows
- add automated approval gates
- create drift detection or policy checks

## Stretch goal
- implement blue/green or slot-based deployment
- add release tags and environment promotion
- create a rollback checklist

## Test plan
- tests fail when code quality is poor
- invalid infra definition blocks deployment
- dev deployment succeeds
- production deployment requires explicit approval or guardrails

### Negative tests and failure cases
- broken YAML workflow
- invalid resource configuration
- secret unavailable at runtime
- app deployed without passing smoke checks

## Security and privacy considerations
- never store secrets in repo files
- use environment variables or secret stores
- limit deployment identity permissions
- review exposure of infrastructure configuration

## Cost risks, usage limits, and teardown checklist
- CI pipelines can consume compute minutes
- infrastructure left alive after testing creates cost
- duplicate dev environments add cost
- clean up after use

## AI-agent prompts
- “Create a minimal GitHub Action pipeline for deployment to Azure with lint, tests, and a smoke test.”
- “Explain how infrastructure as code reduces drift and operational inconsistency compared with portal changes.”

## Questions participants must answer
- Why is infrastructure as code safer than manual portal changes?
- What is the value of a smoke test after deployment?
- How do approval gates reduce risk?

## Deliverables for the repository
- CI workflow files
- infrastructure definitions
- README with deployment flow
- environment variables or config notes
- cleanup instructions

## Reflection exercise and retrospective questions
- What would change if this app were used by real paying customers?
- How would you handle rollback and recovery in production?

## Rubric
- automation quality: 30%
- versioned infra quality: 25%
- security and pipeline hygiene: 20%
- deployment safety: 15%
- documentation: 10%

## Hint 1
Treat deployment automation as part of the application, not a separate project step.

## Hint 2
Always validate the app after deployment, not just before.

## Hint 3
Start with one environment and one deployment stage before adding production gates.
