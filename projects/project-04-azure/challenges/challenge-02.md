# Challenge 2: Build a cloud API using serverless compute, HTTP triggers, configuration, and persistence

## Difficulty
Beginner–Intermediate

## Problem statement and user story
The venue booking platform needs a lightweight API to calculate prices, create bookings, and retrieve venue availability. The API must handle bursts of traffic and avoid unnecessary server management.

## Practical outcome
Participants build and deploy a serverless API that responds to HTTP requests and stores booking data.

## Learning objectives
- understand Azure Functions HTTP triggers
- use application configuration instead of hardcoded values
- connect compute to data storage
- explain why serverless can be a good fit for APIs

## Prerequisites and setup
- basic understanding of HTTP APIs
- local runtime environment for the chosen language
- Azure Functions or equivalent hosting knowledge
- access to a storage account or local data storage for testing

## Azure services under consideration
- Azure Functions
- Azure Storage
- Azure Monitor
- Application settings

### Verified free-tier / cost notes
- Functions pricing depends on plan and usage patterns.
- Storage and monitoring also have usage-based costs.
- Always verify current pricing and free-offer eligibility before deployment.

## Reference architecture and request/data flow
```text
Client -> HTTP trigger -> validation -> storage or database -> response
```

## Functional and non-functional requirements
Functional:
- API accepts booking requests
- API validates required fields
- API returns status or confirmation
- data is persisted

Non-functional:
- small code footprint
- stateless design where possible
- reasonable response times
- logs for failed requests

## Acceptance criteria
- HTTP endpoint works in Azure or via emulator
- booking is created and retrievable
- invalid payloads return clear errors
- configuration is externalized

## Minimal implementation path
1. create function app
2. add HTTP endpoint
3. validate booking payload
4. persist data to storage or a lightweight database
5. return success response

## More advanced implementation path
- add pagination for bookings
- add filtering by date or venue
- move data to a more structured database
- add API versioning

## Stretch goal
- support idempotent booking creation for repeated submissions
- add caching for venue availability
- add a scheduled cleanup job for stale reservations

## Test plan
- valid booking succeeds
- missing fields fail with 400
- duplicate request is handled safely
- data is stored and returned correctly

### Negative tests and failure cases
- invalid date or venue ID
- storage unavailable
- function app misconfigured
- output payload format mismatch

## Security and privacy considerations
- never hardcode secrets in code
- validate all user input
- log safely and avoid exposing private data
- restrict access to admin or internal endpoints

## Cost risks, usage limits, and teardown checklist
- function execution count can grow quickly
- storage operations can add up
- logs and monitoring can create cost over time
- delete all resources after testing

## AI-agent prompts
- “Design a minimal Azure Functions API for a venue booking system. Explain the benefits and drawbacks compared with a continuously running web app.”
- “What application settings should be externalized for a serverless booking API?”

## Questions participants must answer
- Why is a serverless API a good fit for small bursts of traffic?
- What happens when the function is cold-started?
- Why is configuration externalization important?

## Deliverables for the repository
- API source code
- README with endpoint examples
- tests for valid and invalid inputs
- architecture notes
- cleanup instructions

## Reflection exercise and retrospective questions
- When would you choose a traditional API host instead?
- Is serverless always the right solution for a public API?

## Rubric
- API correctness: 30%
- security/config handling: 20%
- reasoning about serverless trade-offs: 25%
- testing quality: 15%
- docs and cleanup: 10%

## Hint 1
Think about the system boundary between user request and persistent data.

## Hint 2
Start with one or two endpoints before designing a broad API surface.

## Hint 3
Use a single payload model and validate it strictly before you add more features.
