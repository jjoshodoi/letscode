# Project 2 — The Slow API

An intentionally working but poorly scaling Product Search API. The API returns
the right answers and its tests pass. Your job is to investigate why it is slow,
improve it, and prove that the improvement is real.

This is a **legacy/performance challenge**, not a feature race. Do not begin by
guessing at optimisations:

> **Measure → Hypothesise → Change → Measure**

## Residency fit

This is the second individual project in the AI Engineering Residency:

1. Project 1 — **Build a Backend** (the Go URL shortener)
2. Project 2 — **Investigate a Backend** (this project)

It follows the residency process:
**Understand → Question → Explore → Design → Build → Test → Measure → Explain → Reflect**.
The code is evidence of engineering judgement, not the only outcome.

Everyone uses the same starter, but chooses a route appropriate to their
starting point. There is no requirement that everyone makes the same changes.

## What is provided

- Go HTTP API backed by SQLite
- Seed data (1,000 products in the application, expandable for experiments)
- Passing endpoint tests
- A Go benchmark
- Docker setup
- CI running tests, `go vet`, and the benchmark

The starter deliberately contains more work than necessary in its request path.
Treat the implementation as an unfamiliar system: form your own hypotheses
from measurements rather than searching for an answer key.

## Prerequisites

- Go 1.20 or newer
- Docker (optional)
- `curl`

The SQLite driver uses CGO. On macOS, install Xcode Command Line Tools if Go
cannot compile the driver. Docker provides a consistent alternative.

## Quickstart

From this directory:

```sh
go mod download
make test
make run
```

In another terminal:

```sh
curl http://localhost:8080/products
curl 'http://localhost:8080/products?category=books'
curl 'http://localhost:8080/products?search=learning'
curl http://localhost:8080/products/1
```

Configuration is optional:

```sh
DB_PATH=data/products.db ADDR=:8080 make run
```

To run with Docker:

```sh
docker build -t slow-api .
docker run --rm -p 8080:8080 -v "$PWD/data:/app/data" slow-api
```

## The challenge

The API has recently started experiencing performance problems as the amount of
data has grown. Investigate the system, identify bottlenecks, and improve its
performance. You must demonstrate that your changes made it faster without
breaking correctness.

Do not optimise blindly. For every performance change, record:

### Before

- What user-visible or system metric are you measuring?
- What is the baseline under a stated workload?
- What evidence suggests this is the bottleneck?

### Change

- What did you change?
- Why did you expect it to help?
- What trade-offs did it introduce?

### After

- What is the new result, using the same workload?
- How much did it improve?
- Did latency, throughput, memory, database size, or correctness get worse?

Use an ADR, a short investigation document, or your PR description for this
record. Documentation is part of the deliverable.

## Routes

| Method | Route | Purpose |
| --- | --- | --- |
| GET | `/products` | List products |
| GET | `/products?category=books` | Filter by category |
| GET | `/products?search=learning` | Search name and description |
| GET | `/products/:id` | Fetch one product |

## Learning routes

### Supported route — guided investigation

1. Run the application and verify the endpoints.
2. Read the tests and implementation before changing code.
3. Define one workload, such as 100 search requests against 1,000 products.
4. Record latency and throughput with a repeatable command or benchmark.
5. Inspect SQL, logs, CPU, memory, and query plans to find where time goes.
6. Write a hypothesis before editing.
7. Make one focused change.
8. Run the tests and repeat the exact baseline workload.
9. Compare the results and record the trade-offs.
10. Repeat only when new evidence justifies another change.

### Independent route

Choose the workload and tools yourself. Establish a baseline, investigate at
least one bottleneck, make a measured improvement, and explain why alternative
solutions were not chosen. Preserve or extend the tests as needed.

Useful starting commands:

```sh
make benchmark
go test ./... -count=1
go test ./internal/api -run '^$' -bench BenchmarkListProducts -benchmem -benchtime=5s
```

### Stretch route

Investigate several layers, such as query construction, indexes, allocation,
concurrency, connection handling, caching, response size, and observability.
Compare at least two approaches and explain their operational costs. Add a
repeatable performance gate to CI with a threshold you can defend rather than
an arbitrary number.

## 100× Data Challenge

The starter contains 1,000 products. Repeat your investigation with 100,000
and then (if your machine permits) 1,000,000 or 10,000,000 records.

Ask:

- What breaks first, and why?
- Does the optimisation still work at a larger scale?
- What becomes the new bottleneck?
- Which metric would you monitor in production?
- Would you change the schema or architecture?
- What would you deliberately not change yet?

Be explicit about hardware, dataset size, request mix, concurrency, and warm-up
when comparing results. A benchmark number without its conditions is not useful
evidence.

## CI and engineering quality

Every push and pull request runs:

```text
Code → Tests → Vet → Benchmark
```

Keep the API behaviour correct. Add tests for behaviour affected by an
optimisation. Keep changes reviewable, and use the project review questions:

1. What problem did we solve and what assumptions did we make?
2. Why this change, and what alternatives existed?
3. What could fail in production?
4. How would this perform, scale, deploy, and operate?
5. What did AI suggest, and what did we verify ourselves?
6. What did we learn?

## AI use

AI may help explain profiling output, suggest hypotheses, generate benchmark
scaffolding, review a query, or challenge trade-offs. It cannot replace the
investigation. If AI produces a change, you must understand it, test it, and
be able to explain why it is appropriate.

## Definition of done

- The original tests still pass.
- At least one bottleneck is supported by evidence.
- At least one change is measured before and after under the same workload.
- Correctness and relevant regressions are checked.
- Trade-offs and rejected alternatives are documented.
- The result is explainable to the group during review.
