# Azure Residency Answers and Facilitator Notes

This file is intentionally separate from the workshop guide so that participants can read the guide without seeing the full answers up front. It is meant to support facilitators, mentors, and learners who want a clear answer key after the workshop.

---

## Part 1: Core definitions and explanations

### What is cloud computing?

Cloud computing is the on-demand delivery of compute, storage, networking, and managed services over the internet. Instead of buying servers and maintaining them on-premises, teams use a provider’s infrastructure and managed services.

Key strengths:

- faster setup time
- easier scaling
- lower operational overhead for common concerns
- usage-based pricing

Key risks:

- hidden cost if services are left running
- dependency on provider services and quotas
- need for monitoring, security, and architecture discipline

### What is a subscription?

A subscription is the billing and access boundary for Azure usage. It sits above resource groups and resources. It usually includes:

- a billing account or contract
- a set of resource groups and resources
- quotas and limits tied to the account
- access identity and permissions

### What is a resource group?

A resource group is the logical container for Azure resources. It is useful for grouping multiple related resources and managing them together.

Example:

- a dev environment resource group for a booking app
- a production environment resource group
- a shared monitoring group

### What is the Azure resource hierarchy?

The common hierarchy is:

- management group (optional)
- subscription
- resource group
- resource

This matters because policies, permissions, governance, and lifecycle are usually managed at different levels. A team can apply a policy at the subscription level and manage day-to-day operations within a resource group.

### What are regions?

A region is a geographic location containing Azure data centers. Regions matter because:

- latency depends on user proximity
- service availability differs by region
- data residency rules may require certain locations
- some offerings are not available everywhere

### What is availability?

Availability is the ability of a service to remain accessible and functional. Azure provides redundancy, regions, availability zones, and service SLAs. A resilient system should expect failures and design around them.

Examples:

- one VM failing should not take down the whole app if designed with redundancy
- a queue or event-driven worker can handle traffic spikes without blocking the web API
- a regional failover strategy may be needed for critical workloads

---

## Part 2: Shared responsibility model

### What is the shared responsibility model?

Azure is responsible for protecting the cloud, while the customer is responsible for protecting what is deployed in the cloud.

Azure is responsible for:

- physical datacenter security
- platform infrastructure reliability
- networking fabric and some platform-level controls

The customer is responsible for:

- application design
- identity and access configuration
- secrets management
- data classification and privacy
- proper network, firewall, and role configuration
- monitoring and governance

### Why does this matter?

A managed service still requires good architectural decisions. For example, Azure Functions may manage the platform, but the developer still needs to securely configure triggers, storage, queue usage, networking, and app settings.

---

## Part 3: Consumption-based pricing

### Why can small experiments still cost money?

Even if the starting usage is low, cloud services can still incur costs because Azure charges for:

- compute hours or execution time
- storage capacity and operations
- message volume and queue retention
- outbound data transfer
- monitoring and log ingestion
- premium features or non-free service tiers

### Example cost risks

- an app runs continuously even when traffic is minimal
- functions process a lot of messages and retries create repeated work
- logs are too verbose and accumulate fast
- a dev environment is left running beyond the learning period
- data transfer or egress charges are higher than expected

### Guidance

Always check:

- current Azure pricing pages
- current free service limits
- region and quota restrictions
- whether the service is free for new accounts, always free, or usage-based

---

## Part 4: Workshop questions and answers

### 0–10 minute questions answered

#### Q: What is the difference between a subscription and a resource group?

A subscription is the billing and access boundary. A resource group is a logical container used to organize and manage related resources. A subscription can contain many resource groups.

#### Q: Why do regions matter?

Regions affect latency, availability, compliance, and service support. An app used by people in Europe may perform better in a European region. Some workloads also need to stay within a certain geopolitical boundary.

#### Q: Why is the shared responsibility model important?

Because Azure manages the platform, but your team still owns many application-level security responsibilities. The service itself does not replace secure configuration, access control, or good application design.

#### Q: Why can a small app still cost money?

Because cloud services are usage-based. Compute, storage, messages, monitoring, and outbound transfer can all accumulate costs, even when the load is small.

#### Q: What is the purpose of a resource hierarchy?

It gives structure for governance, permissions, policy, and lifecycle management. It also helps teams reason about environment boundaries such as dev, test, and prod.

---

## Part 5: Common runtime choice examples

### 1. Static/public web app

Use case: a brochure or public venue pages.

Answer:

- low operational complexity
- good for content that does not change often
- often the cheapest and simplest approach

Trade-off:

- not ideal for complex business logic or secure internal workflows

### 2. Web app backend

Use case: a booking app that validates dates, estimates cost, and stores bookings.

Answer:

- keeps the server-side logic in one place
- easier for forms, sessions, and business rules
- simpler mental model for a small project

Trade-off:

- more responsibility than a static site
- may be less cost-efficient for low or bursty workloads

### 3. Serverless API

Use case: price calculation endpoint or booking confirmation endpoint.

Answer:

- event-driven and highly scalable
- useful for API requests that come in bursts
- low startup effort for a focused function

Trade-off:

- cold starts and eventual consistency concerns can matter
- debugging and runtime behavior may feel less predictable at first

### 4. Queue-triggered worker

Use case: sending confirmation emails or auditing bookings after the main booking is accepted.

Answer:

- decouples front-end request from downstream processing
- protects the user-facing path from slow or failing tasks
- supports retries and failure handling

Trade-off:

- adds more moving parts
- requires idempotency and clear failure handling

---

## Part 6: Answers for the Harbor & Hire booking app scenario

### Why this app is a good learning exercise

The booking app naturally contains:

- a public front end
- a booking or pricing API
- storage or persistence of bookings
- queue-triggered notifications or auditing

This gives enough complexity to discuss Azure hosting and architecture without requiring a huge platform.

### What the app flow looks like

1. User chooses a venue or room.
2. User selects a date and time.
3. App calculates the estimate.
4. User confirms booking.
5. Booking is written to a data store.
6. A worker processes follow-up tasks such as email sending or status updates.

### What architecture choices should be considered?

For a simple first version:

- static or managed web front end for public browsing
- backend API or serverless endpoint for booking logic
- storage or database for bookings
- queue for any delayed tasks

For a more advanced version:

- auth and role-based access
- pipeline automation and IaC
- monitoring and alerting
- secure secrets and identity-based access

---

## Part 7: Questions for the 25–45 minute investigation block

### Hosting a web app

#### Q: What is the main difference between static hosting and a dynamic backend?

A static hosting service is ideal for content that doesn't change per user. A dynamic backend is needed when there is business logic, user state, or a database connection.

#### Q: Why would we choose a managed web app over a fully custom server model?

Because managed hosting removes many infrastructure concerns while still allowing application deployment and scaling.

### Persisting and retrieving application data

#### Q: What are the trade-offs between SQL and document storage?

SQL is better when relationships and transactions matter. Document storage is better when the schema is flexible or the data is JSON-like and evolves quickly.

### Executing code in response to HTTP or events

#### Q: Why use serverless or queued workers?

Because they allow workloads to scale independently and isolate slow or failure-prone operations from the user request path.

---

## Part 8: Good answers to the exit ticket

### 1. What is the difference between a subscription and a resource group?

A subscription is the billing and access boundary; a resource group is a container for resources that logically belong together.

### 2. Why do regions matter in Azure?

Regions affect latency, availability, data residency, service support, and disaster recovery planning.

### 3. What does the shared responsibility model mean in practice?

Azure manages the platform, but the team is responsible for identity, access, application design, secrets, configuration, and monitoring.

### 4. Why can a small cloud app still incur charges?

Cloud services are paid by usage, such as storage, compute, monitoring, message processing, and data transfer.

### 5. Which runtime choice would you choose for a booking app and why?

A small but real booking app is usually best with:

- a managed web front end or app host
- a backend or serverless API for pricing and booking logic
- a database or storage layer for bookings
- a queue or async worker for reminders or notifications

This keeps the system simple but still demonstrates the core cloud concepts.

---

## Part 9: Facilitation notes for common misunderstandings

### “Azure handles everything, so we do not need to think about security.”

False. Azure handles the platform; the team still owns app security, identity, secrets, network design, and business logic validation.

### “It is cheap because it is small.”

Small workloads can still trigger storage, monitoring, and service charges. Cost awareness is a design requirement.

### “Serverless is always the best choice.”

Not always. Some workloads benefit from always-on managed hosting, and some apps need predictable performance or simpler debugging.

### “Queuing means the job is guaranteed to succeed.”

No. Queues improve reliability and decoupling, but they require retry, idempotency, and failure handling.

---

## Part 10: Suggested facilitator summary

A good summary statement for the session is:

“The cloud does not remove architecture from the problem. It shifts where the operational work happens. Azure gives us managed infrastructure, but the team is still responsible for how the application behaves, how it scales, how it is secured, and how it is kept cost-aware.”

---

## Part 11: Recommended answer structure for challenge briefs

For each challenge, the participant should be able to answer these questions:

- What problem does the Azure service solve?
- What would happen if we did not use it?
- What alternatives exist?
- Why is the chosen option appropriate?
- What are the cost and quota limitations?
- What security implications exist?
- What risks appear as traffic grows?
- How would we recover from failure?

These are the key cloud engineering questions for the entire residency.
