Project 1 — URL Shortener

This starter implements a minimal URL shortener service in Go with a JSON-backed file store.

Prerequisites

- Go 1.20+ installed (recommended via Homebrew on macOS):

  brew update && brew install go
  go version  # should print go1.20+

- (Optional) Docker to run the service without installing Go.

Quickstart:

1. make run  # requires Go on PATH
2. curl -X POST -H "Content-Type: application/json" -d '{"url":"https://youtube.co.uk"}' http://localhost:8080/shorten
3. curl http://localhost:8080/<code>

If Go is not installed, use Docker:

  docker build -t urlshortener .
  docker run -p 8080:8080 -v $(pwd)/data:/app/data urlshortener

Quick summary:

* Sends an HTTP POST to /shorten with a JSON body: {"url":"https://example.com"} (the -H sets Content-Type; -d makes it a POST).
* The server decodes the JSON, generates a short code, saves code→URL in the file store (data/urls.json), and responds with JSON: e.g. {"code":"Ab1Cd2"} (HTTP 200).
* After that, visiting http://localhost:8080/Ab1Cd2 issues a 302 redirect to the original URL.

SQLite migration

This project now supports a SQLite-backed repository via the same `Store` interface. To switch storage:

1. The default backend is SQLite, so the server will use `data/urls.db` unless you override `STORE_TYPE` or `STORE_DSN`.
2. To force the legacy JSON store, set `STORE_TYPE=file` and optionally `STORE_PATH=data/urls.json`.
3. Start the service without env overrides to create the SQLite table automatically.
4. If you are migrating from JSON, export the existing mappings and insert them into the `short_urls` table:
   - `sqlite3 data/urls.db "CREATE TABLE IF NOT EXISTS short_urls (code TEXT PRIMARY KEY, url TEXT NOT NULL UNIQUE, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);"`
   - `sqlite3 data/urls.db ".mode json"` and import or backfill values from `data/urls.json`.
5. Validate the app by hitting `/shorten` and `/health` and verifying the database contains unique `code` and `url` values.

The repository abstraction is intentionally small and stable:

```go
repo, err := store.NewRepository("sqlite", "data/urls.db")
if err != nil {
    panic(err)
}
```

Swap the implementation without changing handler logic.

Testing:

    go test ./...
    go test -race ./...
    go vet ./...

GitHub Actions runs the standard test and vet commands for pushes and pull requests.

Errors: 400 for bad JSON, 405 if not POST, 500 on save failure. The implementation is simple (no collision handling, no validation).

Notes:
- Persistence is file-backed (data/urls.json) to keep the starter lightweight.
- Replace the store with SQLite or another DB as an exercise.
- See adrs/0001-record-decisions.md for an initial ADR.

Remaining / High-priority work

1) Input validation & URL normalization
- Validate request payloads (presence and JSON schema).
- Normalize URLs: ensure scheme exists (add https:// when absent), parse using net/url, and reject inv§alid or disallowed schemes (e.g., file://).
- Return clear 4xx errors with helpful messages for client mistakes.
- Suggested approach: small helper NormalizeURL(string) (stdlib net/url) plus tests.

2) Collision handling & idempotent shorten
- Provide idempotency: if a URL already exists in the store, return its existing code instead of creating a new one.
- Avoid code collisions: on generated-code collision, retry generation a bounded number of times, then fall back to a deterministic approach (e.g., base62(hash(URL))+counter).
- Consider adding an index or reverse lookup in the Store to support URL->code queries efficiently.

3) Graceful shutdown, logging, config
- Use http.Server with Shutdown(ctx) and trap signals (os/signal.Notify) for graceful shutdown.
- Replace fmt/log prints with a structured logger (std log is acceptable for starter; consider logrus/zerolog later).
- Read configuration from environment variables (PORT, STORE_PATH, LOG_LEVEL) with sensible defaults.

4) Tests and CI
- Add unit tests for handlers and store behavior (including concurrency tests for Save/Lookup).
- Add e2e tests that start the server in test mode and exercise POST /shorten + GET /{code}.
- Add GitHub Actions workflow to run go test ./... and basic linters (go vet).
- Add a concurrency test that calls Save rapidly from many goroutines to ensure no deadlocks.

5) SQLite-backed store
- Implement a Store variant using SQLite with proper schema and unique constraints for codes and (optionally) URL.
- Use transactions for atomic writes and to enforce uniqueness.
- Provide migration notes and a small repository pattern for easy swapping.

6) Observability & Production Readiness
- Add health and readiness endpoints and basic metrics (Prometheus-friendly counters/histograms).
- Add rate-limiting for POST /shorten (simple token-bucket or per-IP limiter) to prevent abuse.
- Improve Dockerfile for production (multi-stage, non-root user) and add deployment notes.

Implementation notes
- Prefer small, focused PRs that add one capability at a time.
- Document architecture-impacting changes with ADRs and add tests alongside behavior changes.
- Keep the Store interface stable to make swapping persistence implementations straightforward.

If helpful, I can scaffold one of these items (validation and idempotency recommended). Which should be done next?

---

Configuration (env)

The server reads configuration from environment variables with sensible defaults:

- PORT — TCP port to listen on. Defaults to :8080. You can provide just the port number (e.g. 8080) or include the leading colon.
- STORE_PATH — Path to the JSON file used by the file-backed store. Defaults to data/urls.json.
- LOG_LEVEL — Logging verbosity: debug, info, warn, error. Defaults to info.

Graceful shutdown

The server uses an http.Server and traps SIGINT/SIGTERM. On the first termination signal it starts a shutdown with a short timeout (5s) to allow in-flight requests to finish. This ensures file-backed persistence has a chance to complete pending writes.
