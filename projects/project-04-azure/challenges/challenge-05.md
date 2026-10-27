# Challenge 5: Automate a business process using scheduled execution or workflow orchestration, external APIs, and idempotent operations

## Difficulty
Intermediate

## Problem statement and user story
A business needs to reconcile booking status, trigger reminder actions, or push updates to a customer system. This work should happen automatically without requiring staff to manually run it every time.

## Practical outcome
Participants build a scheduled or orchestrated workflow that performs repeated work safely and idempotently.

## Learning objectives
- distinguish scheduled jobs from event-driven jobs
- implement orchestrated workflow logic
- handle retries and duplicate requests safely
- explain why idempotency matters in automation

## Prerequisites and setup
- ability to call external APIs or stubbed endpoints
- familiarity with HTTP requests and status codes
- understanding of time-based triggers or workflow logic

## Azure services under consideration
- Azure Functions Timer Trigger
- Azure Logic Apps
- Azure Storage
- API integration services or API gateway patterns

### Verified free-tier / cost notes
- execution cost depends on trigger volume, connector use, and service tier.
- Pricing and free offers must be checked before deploying workflow services.

## Reference architecture and request/data flow
```text
Timer or schedule -> workflow -> external API -> status storage -> retry or completion record
```

## Functional and non-functional requirements
Functional:
- job runs on schedule
- work can be processed reliably
- results are stored and visible
- failed jobs can be retried safely

Non-functional:
- idempotent operations
- logs and state visibility
- graceful recovery from transient failures

## Acceptance criteria
- workflow runs successfully on a defined schedule
- duplicate or repeated execution does not produce duplicate side effects
- failed requests are recorded and retried or marked explicitly
- status is visible in logs or storage

## Minimal implementation path
1. create scheduled trigger or workflow
2. retrieve pending work
3. call external system or internal API
4. record success or failure
5. repeat safely on retries

## More advanced implementation path
- add durable state tracking
- implement a queue between scheduler and executor
- add webhook or event-driven step after external completion

## Stretch goal
- use a dashboard or log summary for workflow status
- add rate limit handling and backoff
- add replay capability for failed jobs

## Test plan
- successful job run
- transient error then retry
- duplicate schedule invocation does not re-submit work
- external API timeout is handled gracefully

### Negative tests and failure cases
- API returns 500 or timeout
- job runs twice in quick succession
- inconsistent status state across storage and workflow logs
- malformed payload or missing required field

## Security and privacy considerations
- protect API keys and tokens
- log only safe metadata, not user-sensitive data
- validate request payloads before making external calls
- review network and access boundaries

## Cost risks, usage limits, and teardown checklist
- scheduled jobs can create repeated cost even when idle
- retries and API calls increase operational cost
- remove workflow resources after the learning exercise

## AI-agent prompts
- “Describe an idempotent reminder workflow for a bookings platform using Azure Functions or Logic Apps.”
- “How should a scheduled job handle repeated execution without creating duplicates?”

## Questions participants must answer
- Why is idempotency central to scheduled work?
- What is the difference between a timer trigger and a queue-driven worker?
- Why do automation jobs need safe failure handling?

## Deliverables for the repository
- workflow definition or function code
- README with schedule and state model
- examples of retried and successful jobs
- cleanup and monitoring notes

## Reflection exercise and retrospective questions
- What would happen if the automation ran more than once per period?
- Which business process should remain asynchronous and which should stay synchronous?

## Rubric
- workflow correctness: 30%
- idempotency and failure handling: 25%
- operational clarity: 20%
- security and data handling: 15%
- docs and testing: 10%

## Hint 1
Start from the exact state transition of a job: pending, executing, succeeded, failed.

## Hint 2
Think carefully about replay safety before adding more workflow steps.

## Hint 3
Keep the first version easy to reason about: one scheduled job, one API call, and one clear outcome.
