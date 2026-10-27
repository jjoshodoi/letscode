# Challenge 8: Production-readiness capstone with observability, resilience, health checks, recovery, security, and cost management

## Difficulty
Advanced

## Problem statement and user story
The venue booking system is now a real product. It must be reliable enough for real customer traffic, with clear monitoring, secure access, safe deployment, and controlled cost.

## Practical outcome
Participants produce a production-ready architecture and operational plan for the booking platform using the full set of Azure concepts covered in the residency.

## Learning objectives
- integrate the earlier challenges into one coherent architecture
- design for resilience and recovery
- add observability and health checks
- reason about security, availability, and cost together

## Prerequisites and setup
- completion of earlier challenges or a comparable baseline architecture
- access to Azure resources for deployment and monitoring
- ability to explain service choices and trade-offs

## Azure services under consideration
- App Service or Azure Functions
- Azure Storage or database service
- Service Bus or queues
- Azure Monitor or Application Insights
- Key Vault and identity services
- optional pipeline and infrastructure tools

### Verified free-tier / cost notes
- many services have consumption-based pricing and quotas.
- No Azure service should be treated as universally free without checking current pricing and eligibility.
- Cost controls and teardown plans are mandatory in the final design.

## Reference architecture and request/data flow
```text
User -> public web front end -> API -> data layer
                                     -> queue worker for async work
                                     -> logs, metrics, traces, alerts
                                     -> identity and secret access
```

## Functional and non-functional requirements
Functional:
- booking flow works end-to-end
- app can be browsed and booked by real users
- background tasks and scheduled processes function reliably

Non-functional:
- health checks and alerting
- graceful degradation under failure
- secure configuration and access
- cost visibility and cleanup process

## Acceptance criteria
- end-to-end system runs successfully
- failures are observable and diagnosable
- the setup includes monitoring and recovery thinking
- security boundaries are explained
- cost controls are documented

## Minimal implementation path
1. take the prior challenge app and integrate it into one coherent system
2. add health checks and logs
3. add identity and secret management
4. add alerting or monitoring dashboards
5. document cost and cleanup plan

## More advanced implementation path
- add multi-region or failover design considerations
- implement queue retries and DLQ analysis
- add deployment approval gates and rollback playbooks
- create runbooks for common incidents

## Stretch goal
- define a disaster recovery strategy
- add a cost anomaly review process
- automate scale and cleanup policy

## Test plan
- health check passes when app is available
- failed dependency or queue message is logged and recoverable
- unauthorized access is blocked
- cost dashboard or budget alerts can detect spikes

### Negative tests and failure cases
- database unavailable
- identity misconfiguration
- function or queue failure
- repeated or malformed messages
- unexpected spike in temp jobs or traffic

## Security and privacy considerations
- restrict secrets and identities
- use least privilege at every layer
- validate all business inputs and external calls
- ensure logs do not expose private data
- keep environment separation clear

## Cost risks, usage limits, and teardown checklist
- logs, monitoring, and data transfer can be hidden cost drivers
- queue retries and repeated API calls increase usage
- unnecessary scale or extra environments raise cost
- always define cleanup steps and cost-review gates

## AI-agent prompts
- “Review this production-ready Azure architecture for observability, resilience, security, and hidden cost risks.”
- “Suggest the highest-risk failure scenarios for a booking platform and explain how to monitor and recover from them.”

## Questions participants must answer
- Which dependency is most critical to customer experience?
- What is the most likely failure mode in the system?
- What would we monitor first?
- Which security boundary is most important and why?

## Deliverables for the repository
- final app and architecture documentation
- deployment and rollback notes
- security and cost review
- ADRs and operational notes
- monitoring or alert definition

## Reflection exercise and retrospective questions
- If this app had a real production incident, what would be the first evidence you would look for?
- What would you change after one month of real usage?

## Rubric
- end-to-end correctness: 25%
- resilience and recovery planning: 20%
- security and least privilege: 20%
- observability and cost awareness: 20%
- explanation and documentation: 15%

## Hint 1
Production readiness is not about adding everything; it is about protecting the critical path.

## Hint 2
Start by identifying what must stay available even when dependencies fail.

## Hint 3
Monitor the user-facing path first, then the background work, then the secondary systems.
