# Challenge 3: Implement asynchronous processing using queues or events, including retries, duplicate delivery, and failure handling

## Difficulty
Intermediate

## Problem statement and user story
Creating a booking triggers several follow-up actions: confirmation email, audit log updates, or downstream processing. These tasks should not block user interaction. The system must be reliable even when one downstream action fails.

## Practical outcome
Participants implement a queue or event-driven processing pipeline with retries and failure handling.

## Learning objectives
- understand async communication and message-driven patterns
- reason about retry and dead-letter handling
- design around duplicate delivery
- explain why decoupling improves reliability and scale

## Prerequisites and setup
- understanding of APIs or functions
- ability to run a message broker or queue locally or in Azure
- basic understanding of retries and failure recovery

## Azure services under consideration
- Azure Service Bus
- Azure Storage Queues
- Azure Functions
- Azure Monitor

### Verified free-tier / cost notes
- messaging services are usually usage-based and may have quotas or tiers before cost appears.
- Check current Azure pricing and account eligibility before deploying a queue service.

## Reference architecture and request/data flow
```text
Booking API -> queue/topic -> worker function -> success or retry
                                          -> dead-letter queue on failure
```

## Functional and non-functional requirements
Functional:
- booking creation sends a message
- message is processed by a worker
- failures are retried or moved to a dead-letter queue

Non-functional:
- no duplicate side effects on repeated processing
- clear observability for failed jobs
- ability to recover manually if required

## Acceptance criteria
- booking request completes without waiting for slow tasks
- queue consumer processes work
- retries happen on transient failures
- duplicate or poison messages are handled safely

## Minimal implementation path
1. create queue or topic
2. emit a message when a booking is created
3. create a consumer function
4. implement retry and dead-letter behavior
5. test duplicate and failed messages

## More advanced implementation path
- add topic subscriptions with different consumers
- use correlation IDs and tracking metadata
- add scheduled retry or delayed delivery

## Stretch goal
- add a dashboard or operational log for failed jobs
- support replay of dead-letter messages
- implement message ordering constraints where needed

## Test plan
- success path with one message
- transient error followed by retry
- poison message goes dead-letter
- duplicate message does not double-book or send duplicate notifications

### Negative tests and failure cases
- worker crashes halfway through processing
- malformed JSON message
- queue or storage temporarily unavailable
- repeated message due to retry or network issue

## Security and privacy considerations
- validate payload content before processing
- avoid logging sensitive data
- use least-privilege access for service-to-service communication
- keep retries idempotent when effects are external

## Cost risks, usage limits, and teardown checklist
- message volume and retention create cost
- retries can multiply the number of events processed
- dead-letter queues can fill with repeated failures
- always clean up queues and deployed resources after learning

## AI-agent prompts
- “Explain how to build a queue-driven workflow for venue bookings while handling duplicate messages safely.”
- “What is the difference between retry, dead-letter, and idempotent processing in a message-based architecture?”

## Questions participants must answer
- Why not do the work inline in the booking request?
- What is a dead-letter queue used for?
- Why does duplicate delivery matter?

## Deliverables for the repository
- producer and consumer code
- README with retry behavior and failure handling
- test scripts for duplicate/failure scenarios
- notes on queue configuration and cleanup

## Reflection exercise and retrospective questions
- What would happen if a worker processed the same message twice?
- How would you recover from a backlog without redoing work incorrectly?

## Rubric
- correctness under failure: 30%
- idempotency and retry design: 25%
- observability: 20%
- architecture reasoning: 15%
- testing quality: 10%

## Hint 1
Think about “at least once” versus “exactly once.”

## Hint 2
Use one worker and one slow downstream step before solving a complex fan-out design.

## Hint 3
Model the message lifecycle: sent, picked up, processed, retried, or dead-lettered.
