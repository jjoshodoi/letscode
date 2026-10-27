# Set Up
## Get Azure free trial
https://signup.azure.com/signup?offer=ms-azr-0044p&appId=102&l=en-gb&ref=www.google.com&hasfullconsent=false

## Azure CLI
https://learn.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest

# Next steps after Azure free-trial signup

If you already have Azure access through the free trial, do this next before starting the challenges:

1. Verify the free trial is active in the Azure Portal.
   - Sign in to https://portal.azure.com/
   - Check your subscription and billing status.
   - Make sure the Azure free trial is the active subscription.
   - Add a billing limit of $1 to be on the safe side

2. Confirm the subscription is usable for learning.
   - Open the Azure Portal.
   - Open Subscriptions and confirm the free trial is listed.
   - If multiple subscriptions exist, select the correct one.

3. Install Azure CLI locally if you want to work from the terminal.

```bash
az version
az login
az account show --output table
az account list --output table
```

4. Create a resource group for the workshop.

```bash
az group create --name rg-harbor-demo --location eastus
az group list --output table
```

5. Check that the region is supported for the resource type you plan to use.
   - Some services are not available in every region.
   - If `eastus` is unavailable, switch to another supported region.

6. Set a cost guardrail immediately.
   - Use Azure Cost Management or budget alerts where possible.
   - Keep resources small and delete them at the end of sessions.
   - Do not leave a resource group running when you are done.

7. Start with the simplest working app, not the full architecture.
   - Use the venue booking app as the learning domain.
   - Keep the first version to: web app + booking logic + storage + optional queue worker.

8. Use the portal and CLI together.
   - Portal: visual inspection and resource review
   - CLI: repeatable setup, clean commands, and good learning practice

## Cost-safe setup checklist

- confirm your free trial is active
- create one small resource group only
- choose a common region like `eastus` unless unavailable
- avoid premium SKUs until you understand the cost model
- remove the resource group when finished

## Why this was missing earlier

This was an oversight in the initial write-up. The project started with the learning objectives and challenge flow, but it did not explicitly capture the real-world setup path for someone who already has free-trial access. We are fixing that by making the first steps visible and practical before the workshop starts.

# Azure Cloud Engineering Residency

This folder contains the workshop guide, the challenge roadmap, and the challenge briefs for the Azure residency. It is designed for a mixed-skill cohort, with a facilitator-led workshop, then a sequence of progressively deeper cloud engineering challenges.

## What this residency is trying to teach

The goal is not to memorize Azure product names. The goal is to build practical understanding of how cloud systems work, why certain Azure services fit certain problems, and how to reason about trade-offs in design, cost, security, and operations.

Participants should be able to explain:

- what cloud computing is
- what a subscription is and how it differs from a resource group
- what a resource hierarchy looks like in Azure
- how regions and availability are chosen
- what the shared responsibility model means
- why small cloud experiments can still cost money
- when to use managed hosting, serverless, storage, and async processing

## Workshop framing and essential definitions

### Cloud computing

Cloud computing is the delivery of computing resources over the internet on demand. Instead of buying and maintaining servers in a physical room, teams consume services such as compute, storage, networking, and managed databases from a provider.

Cloud is useful because it can:

- reduce infrastructure setup time
- improve scalability and elasticity
- simplify operations through managed services
- allow teams to pay for what they use

It is not “magic.” The cloud still requires architecture, design, monitoring, access control, backup strategy, and cost awareness.

### Subscription

A subscription is the billing and access boundary. It is the unit that ties Azure usage to a billing account and an identity. A subscription can contain many resource groups and many resources.

A good mental model:

- account identity and billing = subscription
- running workloads = resources inside subscription
- logical grouping = resource group

### Resource group

A resource group is a logical container for Azure resources. Resources in the same resource group usually share a lifecycle, deployment, permissions boundary, or environment such as development, test, or production.

Example:

- one resource group for dev
- one for staging
- one for production

### Resource hierarchy

A useful Azure hierarchy is:

- management group (optional, for governance across subscriptions)
- subscription
- resource group
- resource

That hierarchy matters because governance, access control, policy, billing, and lifecycle management are enforced at different levels.

### Region

A region is a geographic area containing one or more Azure data centers. Regions matter because:

- latency changes depending on where users and services are located
- compliance requirements may require data to stay in a certain geography
- services can differ by region
- some features are only available in selected regions

### Availability and resiliency

Availability is about service uptime and resilience. Azure uses concepts such as:

- availability zones
- regional redundancy
- fault domains and update domains
- service-level agreements (SLAs)

A well-designed system should think about:

- what happens if one server fails
- what happens if an entire zone fails
- what happens if an entire region becomes unavailable

The point is not to “eliminate all failure,” but to design for graceful degradation and recovery.

### Shared responsibility model

In Azure, responsibilities are split between the cloud provider and the customer.

Azure is responsible for the security of the cloud, such as:

- physical infrastructure
- network hardware
- datacenter security
- base platform services

The customer is responsible for security in the cloud, such as:

- configuring access control
- protecting secrets
- managing application logic and validation
- keeping data classification and privacy in mind
- securing identity, network access, and deployment pipelines

This matters because a managed service reduces operational overhead, but it does not remove all security responsibility from the team using it.

### Consumption-based pricing

Azure uses consumption-based pricing for many services. You pay for what you use, and that usage can be measured in hours, operations, messages, storage volume, API calls, data transfer, and so on.

Why small experiments can still cost money:

- even a small app can incur storage, monitoring, and egress charges
- idle resources are not always free
- a service may cost more than expected if it scales or writes frequently
- a failed pipeline or runaway job can keep consuming resources

This is why cost-awareness is part of architecture, not an afterthought.

---

## Recommended pre-work before the workshop

Participants should complete these steps before the first session:

1. Sign up for a Microsoft Azure account or use a team-provided sandbox environment.
2. Confirm access to a browser-based Azure portal.
3. Install Azure CLI locally if possible:

```bash
az version
```

4. Verify login:

```bash
az login
az account show
```

5. Review the official guidance pages before the workshop:
   - Azure free account: https://azure.microsoft.com/free/
   - Azure pricing: https://azure.microsoft.com/pricing/
   - Azure App Service overview: https://learn.microsoft.com/azure/app-service/overview
   - Azure Functions overview: https://learn.microsoft.com/azure/azure-functions/functions-overview
   - Azure Storage overview: https://learn.microsoft.com/azure/storage/common/storage-account-overview

6. Confirm whether local emulation is acceptable if provisioning is blocked.

---

## The booking app scenario used throughout the workshop

The workshop app will be a small venue or local business booking app called “Harbor & Hire,” a fictional service where a customer can:

- choose a venue or room
- view a date and availability
- estimate the booking cost
- reserve a slot
- see an order confirmation

This is a good cloud learning app because it has:

- a public website for browsing and simple booking
- a backend for pricing logic and reservations
- a data layer for storing bookings and availability
- optional async processing for follow-up notifications or queue-driven actions

This structure lets the workshop naturally explore:

- static/public browsing pages
- hosted backend APIs
- serverless function entry points
- queue-driven workflows
- storage and data persistence
- security and cost concerns

---

## First workshop: 60-minute Azure foundations session

### Participation format

This is a facilitator-led workshop and should be feasible in one hour. It includes:

- shared conceptual teaching
- a small Azure walkthrough together
- a guided investigation in groups
- short presentation and recap

### Facilitator schedule

#### 0–10 minutes: Foundations

Introduce and define the following:

- cloud computing
- subscriptions
- resource groups
- resource hierarchy
- regions
- availability
- shared responsibility model
- consumption-based pricing

Prompt questions:

- What is the difference between a subscription and a resource group?
- Why does a team care about region selection?
- Why do we still need architecture even if Azure manages infrastructure?
- Why can a service appear cheap and still become expensive at scale?

#### 10–25 minutes: Guided practical Azure exercise

Do one of the following depending on account readiness:

A. Preferred path: use the Azure portal and CLI if the account and permissions work.

```bash
az login
az account show --output table
az group create --name rg-harbor-demo --location eastus
az group list --output table
```

Then inspect:

- subscriptions
- resource groups
- resources
- regions
- pricing and cost pages

B. Fallback path: if account setup fails or the user cannot provision resources yet, walk through the Azure portal in read-only mode and inspect example services using public docs.

C. Local fallback path: run a local lightweight app and compare the required architecture to Azure-managed hosting.

#### 25–45 minutes: Investigation groups

Participants should review these service patterns using the booking app scenario:

1. Hosting a web app
2. Persisting and retrieving data
3. Executing code on HTTP or event triggers

Group questions:

- What problem is this service trying to solve?
- What happens if we do not use it?
- What are the alternatives?
- What are the hidden costs or operational responsibilities?
- What would go wrong under load or failure?

#### 45–60 minutes: Reporting and discussion

Each group gives a short explanation of:

- the main service choice
- the trade-offs
- what they would still need to investigate
- what they think is the biggest risk or cost issue

---

## Common runtime choices to show in the workshop

### 1) Static public web app

Use case:
- landing page, brochure site, event page, simple booking form, public information

Architecture:

```text
Browser -> Azure-hosted public web app -> static content
```

Good for:

- fast content delivery
- lightweight apps
- public information or marketing pages

Trade-offs:

- simple and cheap
- not ideal for complex business logic or secure internal workflows

### 2) Web app backend

Use case:
- dynamic booking app with login, pricing logic, validation, and persistence

Architecture:

```text
Browser -> Web App -> Data store / API layer -> storage or database
```

Good for:

- full-stack apps
- business rules and server-side logic
- managed hosting with less infrastructure responsibility

Trade-offs:

- more operational responsibility than static hosting
- higher cost if scaled unnecessarily

### 3) Serverless API

Use case:
- cost-efficient endpoint for pricing calculation or booking confirmation

Architecture:

```text
HTTP request -> Azure Function -> data access -> response
```

Good for:

- event-driven workloads
- sporadic traffic
- lightweight APIs

Trade-offs:

- simpler scaling but can be harder to reason about at runtime
- cold starts may matter for latency-sensitive APIs

### 4) Queue-triggered worker

Use case:
- asynchronous processing after a booking is created

Architecture:

```text
Booking API -> queue/message -> worker -> email, audit log, downstream sync
```

Good for:

- decoupling producer and consumer
- handling spikes without blocking the user request
- retries and dead-letter patterns

Trade-offs:

- more moving parts
- requires handling duplicates and failure states

---

## The workshop app example: Harbor & Hire

This example is intentionally more interesting than a static landing page. It is a small venue booking system that could be used by event planners and local businesses.

### Basic flow

1. User chooses a venue or room.
2. User picks a date and duration.
3. App estimates the price.
4. User books and receives a confirmation.
5. The system stores the booking and optionally triggers an async follow-up action.

### Why this is useful for cloud learning

This flow naturally demonstrates:

- a public web frontend
- a pricing calculation backend
- a data store for bookings
- a queue/worker for notifications or auditing
- Azure-hosted resources with their own cost models

### Example booking form behavior

- date selected
- venue chosen
- estimated total shown before booking
- user confirms
- app records booking
- background task sends a confirmation or schedules reminder

---

## Workshop questions to ask at each stage

### 0–10 minute definitions

- Can you explain the difference between a resource and a resource group?
- Which is more likely to be shared across multiple workloads: subscription or resource group?
- Why does a region matter for both user experience and compliance?

### 10–25 minute exercise

- What did you notice about the Azure portal and Azure CLI?
- What is the difference between a portal-based deployment and a command-line-based deployment?
- Which Azure resources are easiest to inspect without provisioning everything?

### 25–45 minute exploration

- What problem does this service solve that a local script does not?
- What is the cost sensitivity of this architecture?
- What breaks if the service fails or if a dependency is unavailable?

### 45–60 minute recap

- Which design is the simplest for a small team?
- Which design is the most resilient for spikes or failure?
- Which design is most likely to surprise the team with cost or complexity?

---

## Expected outcomes for the first session

By the end of the hour, participants should be able to:

- explain the core Azure resource model
- understand why subscriptions, resource groups, and regions matter
- distinguish between static hosting, backend hosting, and serverless execution
- identify cost drivers and operational responsibilities
- explain the value of messaging and queue-based decoupling

---

## Troubleshooting and resilience notes for facilitators

If Azure account setup fails:

- switch to a portal-only walkthrough
- inspect public docs and pricing pages
- use local app demos to maintain the architectural discussion

If the CLI fails:

- check whether Azure CLI is installed
- confirm login status
- confirm the correct subscription is selected

If there is no region availability:

- try another supported region
- ask participants to reason about why region choice matters

If the environment is blocked by billing verification:

- keep the workshop focused on concepts, reference architecture, and trade-offs

---

## Exit ticket

Ask each participant to answer these questions in a few sentences:

1. What is the difference between a subscription and a resource group?
2. Why do regions matter in Azure?
3. What does the shared responsibility model mean in practice?
4. Why can a small cloud app still incur charges?
5. Which of the common runtime choices would you choose for a small venue booking app and why?

---

## Challenge roadmap

The full challenge series continues under the challenge briefs in this folder.

| Number | Challenge | Difficulty | Core focus |
|---|---|---:|---|
| 1 | Deploy a real application using Azure hosting | Beginner | hosting, resource groups, regions, Git deployment |
| 2 | Build a cloud API with serverless compute | Beginner–Intermediate | HTTP triggers, config, persistence |
| 3 | Implement asynchronous processing with queues/events | Intermediate | retries, duplicate delivery, failure handling |
| 4 | Choose and implement a data layer | Intermediate | SQL vs document, consistency, cost |
| 5 | Automate a business process | Intermediate | scheduled workflows, idempotency, external APIs |
| 6 | CI/CD and infrastructure as code | Intermediate | pipelines, IaC, repeatability |
| 7 | Secure an application with identity and least privilege | Intermediate–Advanced | RBAC, auth, managed identity |
| 8 | Production-readiness capstone | Advanced | observability, resilience, security, cost |

---

## Challenge materials

The following challenge briefs are included in this project folder:

- [challenges/challenge-01.md](challenges/challenge-01.md)
- [challenges/challenge-02.md](challenges/challenge-02.md)
- [challenges/challenge-03.md](challenges/challenge-03.md)
- [challenges/challenge-04.md](challenges/challenge-04.md)
- [challenges/challenge-05.md](challenges/challenge-05.md)
- [challenges/challenge-06.md](challenges/challenge-06.md)
- [challenges/challenge-07.md](challenges/challenge-07.md)
- [challenges/challenge-08.md](challenges/challenge-08.md)

The full answers and facilitator notes are in:

- [answers.md](answers.md)

The reusable team agent prompts live outside this project in:

- [../agents/README.md](../agents/README.md)

---

## Suggested pacing across 3–4 hours

The roadmap can be stretched across sessions such as:

### Session 1: Foundations and hosting workshop
- workshop and app concept
- resource model and Azure basics
- a simple hosting exercise

### Session 2: Serverless API + data persistence
- build booking API
- store records and pricing data
- discuss database choices

### Session 3: Async processing + automation
- add queue or event-driven processing
- schedule jobs and handle retries
- discuss dead letters and idempotency

### Session 4: Security + CI/CD + production readiness
- secure app, deploy pipeline, add monitoring and cost review
- capstone architecture and final demonstration

---

## Facilitator notes for mixed experience levels

### Beginners

Focus on:

- one service per concept
- a simple deployment path
- understanding cost and lifecycle
- avoiding too much product detail

### Intermediate participants

Focus on:

- architecture trade-offs
- data design decisions
- security boundaries
- testing and reliability patterns

### Experienced developers

Push on:

- operational trade-offs
- cost and reliability analysis
- production readiness and resilience
- less “happy path” reasoning and more failure-mode analysis

---

## Good questions to keep revisiting

- What problem does this service solve?
- What happens if we do not use it?
- Is there a cheaper or simpler alternative?
- Who is responsible for what?
- What gets more complex as traffic grows?
- What breaks when the dependency fails?
- What is the real cost signal here?

---

## Final reminder

Participants should not blindly follow the AI. They should validate architecture decisions against official Azure documentation, pricing, and real deployment constraints. The educational goal is understanding first, automation second.
