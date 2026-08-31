Architecture Overview (starter)

- HTTP server (net/http)
- Handlers layer: routing and HTTP concerns
- Repository boundary: stable Store/Repository interface for persistence
- FileStore: JSON-backed file persistence (data/urls.json)
- SQLiteStore: SQLite-backed persistence with unique constraints on code and URL

Repository pattern

Handlers depend on the `store.Store` interface instead of a concrete backend. The repository is selected in `main` via `store.NewRepository(...)`, which keeps swapping storage implementations trivial.

Example:

```go
repo, err := store.NewRepository("sqlite", "data/urls.db")
if err != nil {
    return err
}
```

This pattern makes it easy to migrate from JSON to SQLite without changing HTTP logic.

Next steps for participants:
- Add input validation, rate limiting, auth
- Add metrics, logging, and health checks
- Add CI, Docker-compose, and deployment scripts
