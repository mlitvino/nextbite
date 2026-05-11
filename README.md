# NextBite

NextBite is a Go backend for restaurant discovery and recommendation. It focuses on low-latency ranking, modular scoring, and a clean separation of API, server wiring, and data access.

For detailed architecture and requirements, see [SPEC.md](SPEC.md).

## Goals
- Provide personalized restaurant recommendations.
- Keep latency low with scalable, bounded concurrency.
- Make scoring logic modular and easy to extend.
- Support real-time updates when data changes.

## Current Features
- Gin-based HTTP server.
- Global API prefix `/api`.
- Health endpoint: `GET /api/health`.
- Users API: `GET /api/users`, `POST /api/users`.
- In-memory user store.

## Project Layout
The Go module lives under `backend/`.

```
backend/
	cmd/nextbite/       # app entrypoint
	internal/api/       # routes and handlers
	internal/server/    # HTTP wiring and middleware
	internal/models/    # domain models
	internal/store/     # persistence interfaces and implementations
```

## Build and Run
From the repo root:

```bash
make backend
make run
```

## API
All endpoints are under `/api`.

- `GET /api/health`
	- Returns `{ "status": "ok" }`
- `GET /api/users`
	- Returns `{ "items": [] }`
- `POST /api/users`
	- Body: `{ "name": "...", "username": "..." }`
	- Validates that both fields are non-empty

## Development Notes
- Go module is nested in `backend/`. Run `go` commands from there.
- Use `make tidy-backend` to update dependencies.
