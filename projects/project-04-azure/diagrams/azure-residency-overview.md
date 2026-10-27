# Azure Residency Overview Diagram

## High-level architecture

```text
                 ┌──────────────────────────────┐
                 │        Users / Browser       │
                 └──────────────┬───────────────┘
                                │ HTTPS
                                v
                 ┌──────────────────────────────┐
                 │   Public Web Front End       │
                 │   Azure-hosted website       │
                 └──────────────┬───────────────┘
                                │
                                │ booking / pricing requests
                                v
                 ┌──────────────────────────────┐
                 │   Backend API / Functions     │
                 │   validates bookings          │
                 │   calculates price            │
                 └──────────────┬───────────────┘
                                │
                 ┌──────────────┼───────────────┐
                 │              │               │
                 v              v               v
      ┌────────────────┐  ┌──────────────┐  ┌──────────────────┐
      │ Booking Store  │  │ Config /     │  │ Queue / Worker    │
      │ SQL or Storage │  │ secrets store│  │ notifications / │
      │ or Document DB │  │ Key Vault    │  │ audit / retries   │
      └────────────────┘  └──────────────┘  └──────────────────┘
             │
             │
             v
      ┌──────────────────────────────┐
      │ Monitoring / Logs / Alerts   │
      │ Azure Monitor / App Insights  │
      └──────────────────────────────┘
```

## Key learning story

This diagram is intentionally simple, but it teaches the same concepts as the full Azure architecture:

- public and internal boundaries
- managed hosting versus application logic
- persistence and configuration boundaries
- async work and retries
- monitoring, cost observation, and security

## Why the booking app is a good example

The Harbor & Hire booking flow naturally demonstrates:

- a public front end
- a pricing + validation backend
- a data layer for bookings
- optional queue processing for confirmations or reminders
- security concerns around secret and role configuration
- cost drivers from storage, compute, monitoring, and retries
