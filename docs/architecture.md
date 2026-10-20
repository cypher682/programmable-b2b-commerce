# Architecture

## Service shape (v1)

```
cmd/b2b              HTTP API + lifecycle
internal/contract    versioned contract aggregate, state machine
internal/eval        deterministic rule evaluation
internal/allocation  priority reservation / allocation
internal/workflow    durable saga orchestration
internal/ledger      double-entry postings
internal/receipts    execution receipts
internal/outbox      transactional outbox + dispatcher
```

## Determinism contract

Evaluation must depend only on:

1. Contract version id and its canonicalized terms.
2. Order command payload.
3. Explicitly supplied effective time.
4. Current allocation state read inside the transaction.

No wall clock, randomness, or ambient mutable state inside the evaluator. Every evaluation emits an `ExecutionEvent` containing the inputs hash, the contract version, the chosen branch, and the resulting effects, so a dispute can replay the decision exactly.

## Contract language constraints (v1)

Structured JSON, not Turing-complete:

- pricing tiers by quantity or period volume
- volume commitments with windows
- allocation priority (integer rank + tie-break rule)
- payment terms (net days, currency)
- penalties / rebates as bounded formulas

## Consistency model

| Domain | Guarantee | Mechanism |
|--------|-----------|-----------|
| Contract version pinning | strong | immutable version rows; orders reference a version |
| Priority allocation | strong within a partition | serialized allocation per (tenant, sku, window) |
| Ledger | strong + balanced | single transaction per posting batch |
| Cross-party workflow | eventual with compensation | saga + idempotent steps + outbox |
| Audit / replay | complete | append-only events with input hashes |
