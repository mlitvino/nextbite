# Copilot Instructions

## Project Context
- Project: NextBite recommendation engine in Go.
- Product/architecture spec: See [SPEC.md](../SPEC.md) for goals, APIs, scoring, and constraints.
- Repository summary: See [README.md](../README.md).

## Coding Rules
- Prefer clear, idiomatic Go with small, focused packages.
- Use context-aware APIs for request handling and cancellation.
- Bound concurrency (worker pools) and avoid unbounded goroutines.
- Keep scoring logic modular and testable (interfaces or function types).
- Avoid global mutable state unless required for performance.

## Structure Guidance
- Separate API, scoring, and data access layers.
- Provide a dedicated package for scoring signals.
- Keep data models in a shared package.

## Quality and Testing
- Add unit tests for scoring and aggregation.
- Include benchmarks for hot paths when relevant.
- Prefer deterministic results for the same inputs.

## Style
- Use ASCII only unless needed.
- Add short comments only for non-obvious logic.
- Avoid overly clever abstractions.
