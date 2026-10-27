# Residency Roadmap

## Overview

This roadmap is designed for a mixed-skill group and can be delivered across 3–4 sessions of one hour each.

## Session plan

### Session 1: Cloud foundations and Azure basics

Focus:
- cloud concepts
- subscriptions and resource groups
- regions and availability
- shared responsibility model
- pricing fundamentals
- first Azure portal and CLI exploration

### Session 2: Hosting and backend patterns

Focus:
- public web app hosting
- backend API design
- local vs managed hosting
- data persistence and app configuration

### Session 3: Async processing and business automation

Focus:
- queues and retries
- event-driven patterns
- scheduled jobs and workflow orchestration
- idempotency and failure handling

### Session 4: Security, CI/CD, and production readiness

Focus:
- identity and access
- secrets and managed identity
- deployment automation and IaC
- observability, resilience, cost, and cleanup

---

## Challenge sequence

1. Deploy a real application using Azure hosting, resource groups, regions, and Git-based deployment
2. Build a cloud API using serverless compute, HTTP triggers, configuration, and persistence
3. Implement asynchronous processing using queues or events, including retries, duplicate delivery, and failure handling
4. Choose and implement a data layer, comparing relational and document-based data modelling and cost trade-offs
5. Automate a business process using scheduled execution or workflow orchestration, external APIs, and idempotent operations
6. Create a CI/CD pipeline that tests and deploys the application and introduce infrastructure as code
7. Secure an application using identity, authentication, authorization, least privilege, and managed identities where supported
8. Create a production-readiness capstone with observability, resilience, health checks, recovery, security, and cost management

---

## Example system used throughout

The application domain is a small venue and booking app called Harbor & Hire.

This scenario supports:
- public web experience
- booking and pricing logic
- backend API or serverless function
- persistence for bookings
- background tasks and async processing

This gives the cohort a realistic system without requiring a large platform.
