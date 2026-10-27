# Challenge 1: Deploy a real application using Azure hosting, resource groups, regions, and Git-based deployment

## Difficulty
Beginner

## Problem statement and user story
A small business needs a public website for room or venue bookings. Staff want a simple way to publish availability, quote prices, and accept bookings without managing a server fleet.

## Practical outcome
Participants deploy a working web application to Azure and explain the hosting choice.

## Learning objectives
- understand Azure resource groups and regions
- compare hosting approaches for public apps
- deploy an app using a managed service
- reason about security and cost for an internet-facing application

## Prerequisites and setup
- Azure account or sandbox environment
- GitHub repo with a web app
- Azure CLI installed if possible
- access to Azure portal or a facilitator-provided subscription

## Azure services under consideration
- Azure App Service
- Azure Static Web Apps
- Azure Monitor
- Azure Resource Group

### Verified free-tier / cost notes
- Free and reduced-price options vary by account type, region, and service configuration.
- Do not assume an Azure service is free in all cases.
- Always check the current Azure pricing and free-offer pages before deploying.

## Reference architecture and request/data flow
```text
Browser -> Azure-hosted web app -> app settings / config
        -> optional monitoring / logs
```

## Functional and non-functional requirements
Functional:
- app loads and shows venue or room availability
- users can choose a date
- users can estimate a booking price
- booking is accepted and stored

Non-functional:
- deployable without manual server management
- accessible over HTTPS
- basic monitoring available
- app settings do not live in source control

## Acceptance criteria
- app is deployed to Azure in one environment
- user can browse the app and reserve a slot
- deployment path is documented
- cost and cleanup risks are noted

## Minimal implementation path
1. create resource group
2. provision hosting resource
3. configure application settings
4. deploy app from GitHub or CLI
5. verify route works and log errors

## More advanced implementation path
- add deployment slots or staging environment
- add custom domain and TLS
- add health checks and alerts

## Stretch goal
- add a release pipeline with branch-based deployments
- use environment variables for pricing and feature flags

## Test plan
- app loads without error
- invalid date selection is rejected
- price is calculated correctly
- deployment failure is visible in logs

### Negative tests and failure cases
- missing app setting
- unsupported region
- bad GitHub deployment branch
- app not reachable after deployment

## Security and privacy considerations
- never commit secrets
- restrict admin access
- ensure only required ports and endpoints are exposed
- avoid public debug modes in app settings

## Cost risks, usage limits, and teardown checklist
- keep the app in a low-cost tier during learning
- check monitoring and log retention costs
- delete resource group after the challenge when no longer needed

## AI-agent prompts
- “Compare Azure App Service and Static Web Apps for a small venue booking application. What are the cost and maintenance differences?”
- “Propose a minimal Azure deployment for a web app with one production branch and basic monitoring.”

## Questions participants must answer
- Why is a resource group useful for a web app?
- Why does region choice matter?
- Why is a managed hosting service preferable to manually managing a VM here?

## Deliverables for the repository
- README with architecture and deployment notes
- source code
- infrastructure or deployment notes
- screenshots or URL proof
- cleanup instructions

## Reflection exercise and retrospective questions
- What would change if the app had 10x more users?
- What operational work is still left even with managed hosting?

## Rubric
- deployment success: 25%
- cloud concept understanding: 25%
- cost and cleanup awareness: 20%
- documentation quality: 15%
- security awareness: 15%

## Hint 1
Think about the difference between content-only hosting and business-logic hosting.

## Hint 2
Consider App Service or Static Web Apps rather than a custom VM.

## Hint 3
Start with one app and one deployment target; do not add auth, database, and queue features in version 1.
