# NextBite Recommendation Engine - Technical Spec

## Overview
NextBite is a backend recommendation service that ranks and suggests menu items for users based on lightweight, heuristic signals. The MVP focuses on fast, deterministic ranking and clear extensibility, not ML training.

## Goals
- Deliver personalized item rankings with low latency.
- Keep scoring logic modular so new signals can be added safely.
- Provide a simple candidate generation and ranking pipeline.
- Stay deterministic for the same inputs.

## Non-Goals
- UI or frontend delivery.
- Machine learning model training.
- Real-time streaming updates (batch or request-time computation only).

## Functional Requirements
- Accept authenticated requests and identify the current user via session.
- Generate a candidate set using simple heuristics (popular, recent, category matched).
- Rank candidates using weighted signals.
- Return top-N recommendations with stable ordering rules.
- Expose REST endpoints for recommendations and auth flows.

## Non-Functional Requirements
- P50 latency under 100ms, P95 under 250ms for top-N requests in the MVP.
- Deterministic results for the same inputs.
- Bounded memory growth for in-memory catalogs.
- Graceful behavior when signals are missing.

## Constraints and Limitations
- Data freshness depends on the update strategy (batch or request-time).
- Scoring remains heuristic unless ML services are introduced.
- Large catalogs may require pagination or external storage.

## Architecture
### High-Level Components
- API Layer: Gin handlers for request parsing and response shaping.
- Recommendation Service: candidate generation + scoring + ranking.
- Store Layer: data access for items, users, and interactions.

### Data Flow
1. Receive request with user session.
2. Build candidate set (union of popular, recent, category matched, or user history).
3. Score candidates with weighted signals.
4. Sort and return top-N with stable tiebreakers.

## Scoring and Ranking
### Initial Signals (MVP)
- Popularity (orders or views over a recent window).
- Category affinity (user preference inferred from history).
- Freshness (newer items are boosted).

### Aggregation
- Weighted sum with configurable weights.
- Stable tiebreakers: score, popularity, item ID.

## API Specification (Current + Planned)
- POST /api/auth/signup
  - Body: {name, username, password}
  - Response: created user
- POST /api/auth/login
  - Body: {username, password}
  - Response: user + session cookie
- POST /api/auth/logout
  - Response: {status: "ok"}
- GET /api/me
  - Response: current user
- GET /api/recommendations (planned)
  - Query params: limit, context
  - Response: list of items with scores and metadata

## Data Model (Logical)
- User: id, name, username
- Item: id, name, category, price, created_at
- Interaction: user_id, item_id, type, ts
- ItemStats (optional): item_id, orders_7d, views_7d

## Possible Implementations
### In-Memory MVP
- Load items and recent stats into memory.
- Generate candidates from simple lists.
- Score and sort in memory.

### Store-Backed
- Move items/interactions to a database.
- Cache hot stats in memory.

## Tools and Dependencies
- Go 1.22+
- Gin for HTTP routing
- context, sync for request cancellation and concurrency control

## Testing Strategy
- Unit tests for each signal and aggregation.
- Integration tests for recommendation endpoint.

## Open Questions
- What items and interactions are available initially?
- What weights should be used in the first release?
- How large is the expected item catalog?
