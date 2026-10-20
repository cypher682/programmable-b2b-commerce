# Programmable B2B Commerce & Contract Execution Network

Multi-tenant platform where buyers and sellers define commercial contracts — volume commitments, tiered pricing, allocation priority, quality gates, payment terms, penalties — that execute as durable, auditable workflows across organizational boundaries.

## Status

Scaffold. Contract DSL, evaluation engine, allocation, and saga orchestration are not implemented yet.

| Stage | Scope | State |
|-------|-------|-------|
| 0 | Service skeleton, config, health, CI-ready build | done |
| 1 | Contract model, versioning, state machine | next |
| 2 | Deterministic pricing + volume commitment evaluation | planned |
| 3 | Reservation + priority allocation under concurrency | planned |
| 4 | Durable workflow / saga layer | planned |
| 5 | Double-entry ledger integration | planned |
| 6 | Execution receipts + audit queries | planned |
| 7 | Penalties, rebates, credits | planned |
| 8 | Multi-party simulation + hardening | planned |

## Documents

- [Architecture](docs/architecture.md)

## Quick start

```sh
go test ./...
go run ./cmd/b2b
curl localhost:8081/healthz
```

Configuration is environment-based; see `.env.example`.

## Layout

```
cmd/b2b            service entrypoint
internal/config    env config
internal/health    liveness/readiness
docs               architecture and decisions
deploy             local dependencies (Postgres)
migrations         schema migrations
```
