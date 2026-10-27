# Shared Team Agents

This folder is intentionally outside the project folders so that reusable AI agent guidance can be shared across multiple challenge projects without being buried inside one application repository.

## Purpose

These agent instructions are meant to help the team:

- reuse good patterns across workshops and projects
- standardize cloud architecture prompts
- keep AI-assisted work grounded in official Azure documentation
- reduce hallucination by requiring evidence before implementation

## Suggested layout

- `azure-guidance.md`: general Azure and cloud design prompts
- `architecture-review.md`: prompts for reviewing proposed designs
- `code-review.md`: prompts for reviewing generated code and tests
- `security-review.md`: prompts for access control, secrets, and identity checks
- `cost-review.md`: prompts for checking usage, quotas, and cost-risk trade-offs

## Team guidance

AI should be used to help with:

- requirements clarification
- architecture alternatives
- implementation planning
- documentation review
- testing ideas
- security review
- deployment troubleshooting
- cost and resilience review

AI should not be treated as the final authority. The human remains responsible for:

- architecture choices
- service selection
- security decisions
- test quality and evidence
- final explanation of the design

## Core rule for all agents

Before accepting a solution, the participant should verify it against official Azure documentation and their own understanding of the system.

The agent should ask clarifying questions first, identify uncertainties, and only then suggest implementations.

## Example reusable agent prompt

> You are a senior Azure cloud architect helping a mixed-skill team. Explain the trade-offs between App Service, Azure Functions, and a queue-driven worker for a venue booking app. Use only authoritative Azure documentation as the reference source. Before proposing a solution, ask 3–5 important clarifying questions and highlight assumptions. Then propose a minimal architecture and explain cost, security, and operational responsibilities.

This repo is intended to grow over time as the team builds its own preferred prompt library.
