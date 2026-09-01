# Project 2 — The Slow API

This project is designed for a 60-minute live session where everyone reads the
same tiny codebase, figures out what's slow, and improves it together.

The goal is not to write a large application. The goal is to make a small API
obvious enough that a group can understand it quickly, find the bottlenecks, and
try real improvements without getting lost in complexity.

## Quickstart

```sh
go mod download
make test
make run
```

Example requests:

```sh
curl http://localhost:8080/products
curl 'http://localhost:8080/products?category=books'
curl 'http://localhost:8080/products?search=learning'
curl http://localhost:8080/products/1
```

Optional config:

```sh
DB_PATH=data/products.db ADDR=:8080 make run
```

And for benchmarking:

```sh
make benchmark
go test ./... -count=1
go test ./internal/api -run '^$' -bench BenchmarkListProducts -benchmem -benchtime=5s
```

With Docker:

```sh
docker build -t slow-api .
docker run --rm -p 8080:8080 -v "$PWD/data:/app/data" slow-api
```


## Little guidance

This is intentionally small, fast to read, and easy to reason about.

Start by reading these files:

- `internal/api/api.go`
- `internal/store/store.go`
- `internal/api/api_test.go`

The app is a Go HTTP API backed by SQLite. It has a few routes and a small
amount of business logic. That is the point: the performance problems should be
fairly obvious once you look at the query path.

Before changing anything, ask:

- What happens when `/products?search=learning` runs?
- Are we doing unnecessary work for every product in the database?
- Are we issuing extra queries inside a loop?
- Are we searching/filtering in Go instead of in SQL?
- Are there missing indexes or repeated work that would scale badly?

A good first pass is to look for anything that scales with the number of rows,
especially in `FindProducts` and `FindProduct`.

If you want the minimal version of the challenge, here are the likely hotspots:

- Selecting every row from `products` even when a category is supplied
- Running `COUNT(*)` for each product in a loop
- Doing case-insensitive string matching in Go over every product row
- No database indexes on the columns used for filtering

This is enough to make the session interesting without requiring a large codebase
or a lot of setup.

## Medium guidance

This is the format for a live tuning session:

1. Run the app and hit the API.
2. Read the two core files: the handler and the store layer.
3. Measure the baseline with the existing benchmark.
4. Identify the obvious bottleneck and write down a hypothesis.
5. Change one thing.
6. Repeat the benchmark under the same conditions.
7. Compare results and explain why the change helped.

Useful commands:

```sh
go mod download
make test
make run
```

Then in another terminal:

```sh
curl http://localhost:8080/products
curl 'http://localhost:8080/products?category=books'
curl 'http://localhost:8080/products?search=learning'
curl http://localhost:8080/products/1
```

And for benchmarking:

```sh
make benchmark
go test ./... -count=1
go test ./internal/api -run '^$' -bench BenchmarkListProducts -benchmem -benchtime=5s
```

The obvious area to investigate is `internal/store/store.go`.

Look for these patterns:

- `SELECT id, name, description, category, price, stock FROM products` with a full scan
- `for rows.Next()` followed by `strings.Contains(...)`
- `db.QueryRow("SELECT COUNT(*) FROM reviews WHERE product_id = ?")` inside a loop
- no `INDEX` statements for category or product lookup

That is the heart of the challenge. The session should feel like a slow, focused
engineering review: read the code, spot the bottleneck, make a measured fix, and
explain the trade-off.

## A lot of guidance

If you want deeper direction before the session, this is the path to follow.

### Where the bottlenecks are

The code is intentionally not trying to hide the issue. The slow path is in the
store layer:

- `FindProducts` fetches all products from the table and filters in Go.
- `strings.Contains(strings.ToLower(...))` runs on every row for every search.
- `FindProducts` also issues one review count query per product.
- The database has no indexes for common access patterns like category filtering
  or product ID lookup as the dataset grows.

This means the problem becomes worse as the dataset grows, and the bottleneck is
very easy to demonstrate.

### Good improvement ideas

A few realistic improvements are obvious and worth trying:

1. Move filtering into SQL.
   - Filter by category in the query instead of doing it in Go.
   - Use SQL `WHERE` clauses for exact and partial matches where appropriate.

2. Avoid the N+1 query problem.
   - Instead of querying review counts one product at a time, aggregate counts in a
     single query using `LEFT JOIN` and `GROUP BY`.

3. Add indexes.
   - Index `products(category)`
   - Index `products(id)` (already primary key, so this is effectively there)
   - Consider a full-text or LIKE-friendly index if search becomes prominent

4. Reduce work in the hot path.
   - Fetch only the fields required for the endpoint.
   - Stop building large in-memory result sets when the request can be restricted.

5. Benchmark before and after.
   - Measure the same search workload before and after each change.
   - Keep the workload consistent so the comparison is fair.

### Example SQL direction

These are exactly the kinds of optimisations to explore:

```sql
CREATE INDEX idx_products_category ON products(category);
```

And instead of:

```go
for rows.Next() {
    // filter in Go
}
```

prefer:

```sql
SELECT ... FROM products WHERE category = ?
```

And for reviews:

```sql
SELECT p.id, COUNT(r.id) as review_count
FROM products p
LEFT JOIN reviews r ON r.product_id = p.id
GROUP BY p.id;
```

### Suggested live session flow

Use this as a 60-minute plan:

- 10 minutes: read the API and store layer
- 10 minutes: run the benchmark and inspect the hot path
- 15 minutes: form a hypothesis and patch one bottleneck
- 10 minutes: rerun the benchmark and compare results
- 10 minutes: discuss trade-offs, rejected ideas, and what to improve next

### Definition of success

The session succeeds if the group can clearly answer:

- What is slow?
- Why is it slow?
- What did we change?
- What changed in the numbers?
- What are the trade-offs?
- What would we do next if we had 30 more minutes?

The beauty of this project is that the bottlenecks are not hidden. The code is
small enough to read quickly, but the optimization opportunities are strong
enough to drive a real engineering discussion.
