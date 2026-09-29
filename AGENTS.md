# AGENTS.md

## Goal

Build production-ready Go code that is simple, readable, maintainable, and feature complete without unnecessary abstraction.

## Core Principles

- Prefer clear, idiomatic Go over clever code.
- Keep solutions simple, explicit, and easy to reason about.
- Avoid premature abstraction, generic frameworks, and over-engineering.
- Write code that is safe for production, not just enough to pass tests.
- Optimize for correctness, clarity, and long-term maintainability.

## Go Practices

- Follow standard Go formatting with `gofmt` and `go vet`.
- Use small interfaces defined by the consumer, not the producer.
- Keep functions short and focused.
- Return errors explicitly and wrap them with useful context.
- Do not panic for normal application errors.
- Prefer context-aware APIs for I/O, database, HTTP, and external calls.
- Avoid package-level mutable state unless there is a strong reason.
- Use dependency injection through constructors and explicit parameters.
- Keep exported names documented when they are part of the public API.

## Error Handling

- Handle every error intentionally.
- Use sentinel errors only when callers need to compare them.
- Prefer `errors.Is` and `errors.As` for error checks.
- Do not hide errors behind vague messages.
- Log errors at the boundary where they are handled, not at every layer.

## Testing

- Add tests for meaningful behavior, edge cases, and error paths.
- Prefer table-driven tests when they improve clarity.
- Avoid brittle tests tied to implementation details.
- Use mocks or fakes only when they make the test simpler.
- Keep tests fast, deterministic, and easy to read.

## Concurrency

- Use goroutines only when concurrency is necessary.
- Always define ownership, cancellation, and cleanup behavior.
- Pass `context.Context` where cancellation or deadlines matter.
- Avoid data races and shared mutable state.
- Use channels for coordination, not as a replacement for simple function calls.

## APIs and Boundaries

- Validate inputs at system boundaries.
- Keep transport concerns separate from business logic.
- Do not leak database models directly into API responses unless intentionally designed.
- Keep request and response types explicit.
- Maintain backward compatibility unless a breaking change is required and documented.
- When changing any API contract, update `docs/openapi.yaml` in the same change.

## Feature Growth and Responsibility Boundaries

- Before extending a feature, inspect its existing responsibilities and dependencies. Do not automatically append new behavior to service.go or repository.go.
- Organize implementations around cohesive use cases and transaction ownership. Separate responsibilities that require different dependencies or change for different reasons.
- Keep dependencies narrow and consumer-owned. Do not expand a shared interface merely to accommodate an unrelated use case. Treat production Go files exceeding roughly 500 lines as a review trigger, not a hard limit. Before growing them further, extract the responsibility being changed or explain why keeping it together is clearer. Generated files, migrations, and tests require separate judgment.
- Moving methods into multiple files is not sufficient if one type still owns unrelated responsibilities. Split ownership where appropriate.
- Keep composition facades thin: wiring and delegation, not business logic. Preserve transaction boundaries, locking, isolation, authorization, error behavior, and public contracts when extracting responsibilities. Never split an atomic operation into independently committed writes.
- Avoid generic repository frameworks, speculative abstractions, and one-method types created only to satisfy a size target. Keep refactoring proportional to the requested change. Request approval before undertaking broader architectural cleanup.
- Verify extracted behavior and failure paths with meaningful tests.

## Database

- Keep transactions small and explicit.
- Handle rollback paths correctly.
- Avoid hidden database calls in unexpected places.
- Make migrations safe, reversible when practical, and easy to review.
- Use clear constraints and indexes instead of relying only on application logic.

## Security

- Never hardcode secrets, credentials, tokens, or private keys.
- Validate and sanitize external input.
- Use least-privilege access for external services and databases.
- Avoid logging sensitive data.
- Treat authentication, authorization, payments, and user data as high-risk areas.

## Before Finishing

- Run formatting, linting, and tests when available.
- Remove dead code, debug logs, and unused dependencies.
- Check edge cases and failure paths.
- Keep the final change focused on the requested task.
