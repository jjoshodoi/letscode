# Challenge 4: Choose and implement a data layer, comparing relational and document-based data modelling, query patterns, consistency, and cost

## Difficulty
Intermediate

## Problem statement and user story
The venue booking app needs to store bookings, venue availability, and customer details. Some operations are highly relational, while others vary by venue and date. The team needs a data design that fits real access patterns and cost constraints.

## Practical outcome
Participants compare a relational model and a document model, then choose and justify a data design for the app.

## Learning objectives
- compare relational and document data models
- identify query patterns, consistency needs, and cost drivers
- explain why not all data needs the same database design
- justify a data architecture with evidence

## Prerequisites and setup
- basic database understanding
- comfort with JSON or structured records
- knowledge of filtering and query patterns

## Azure services under consideration
- Azure SQL Database
- Azure Cosmos DB
- Azure Storage for files or unstructured content

### Verified free-tier / cost notes
- database and storage services have usage-based cost models and quotas.
- Current pricing and free eligibility should be checked before deployment.

## Reference architecture and request/data flow
```text
Frontend -> API -> relational DB or document DB -> query results
                       \-> optional storage for files/attachments
```

## Functional and non-functional requirements
Functional:
- store booking data
- filter by venue and date
- support updates to booking status
- retrieve pricing and summary records

Non-functional:
- efficient reads for common queries
- cost-aware data model
- support for growth without redesigning everything

## Acceptance criteria
- chosen data design matches the app’s needs
- main query patterns are explained
- at least one alternative design is discussed
- security and cost implications are noted

## Minimal implementation path
1. choose relational or document model
2. create schema or document structure
3. add sample records
4. show queries for booking lookup and summary generation

## More advanced implementation path
- compare a second data model in a design note
- add indexing or partitioning strategies
- model a hybrid solution using SQL for transactions and blob storage for files

## Stretch goal
- build a small benchmark or comparison document showing usage patterns and trade-offs
- add a reservations summary query or dashboard

## Test plan
- create booking
- find booking by ID
- list bookings for a date
- update booking status
- invalid or missing records return clean errors

### Negative tests and failure cases
- malformed data input
- duplicate booking reads
- incorrect date or venue query
- large record growth with poor indexing

## Security and privacy considerations
- classify sensitive data fields
- restrict access to admin or service identities
- encrypt data in transit and at rest
- avoid storing secrets in records or logs

## Cost risks, usage limits, and teardown checklist
- reads, writes, and storage volume affect cost
- database throughput and redundancy increase cost
- unused test data should be cleaned up
- confirm teardown plan before deployment

## AI-agent prompts
- “Compare Azure SQL and Cosmos DB for a small venue booking domain. Which is the better fit for a date-based booking flow?”
- “What are the consequences of choosing the wrong data model for a growing reservation platform?”

## Questions participants must answer
- When is a relational model better than a document model?
- What query patterns make a document store attractive?
- Why do cost and consistency matters impact the database decision?

## Deliverables for the repository
- schema or data model document
- sample queries
- README with design rationale
- ADR or design note

## Reflection exercise and retrospective questions
- What would change if the app became a multi-location chain?
- Which part of the app should remain relational even if the rest is document-based?

## Rubric
- model fit: 30%
- explanation of trade-offs: 25%
- query and consistency reasoning: 20%
- cost awareness: 15%
- documentation quality: 10%

## Hint 1
Think about the real query patterns before choosing the database.

## Hint 2
Relational data is usually better for clear relationships and transactions; document data is often better for flexible JSON-like records.

## Hint 3
Do not optimize for every possible future case in version 1.
