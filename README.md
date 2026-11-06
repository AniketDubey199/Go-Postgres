# Go-Postgres

A production-ready PostgreSQL integration module written in Go. This repository demonstrates robust patterns for establishing and managing database connections, configuring connection pools for high concurrency, handling transactions safely, and executing core SQL operations cleanly and idiomatically.

---

## Features

- **Production-Grade Connection Pooling:** Optimized configuration for idle connections, maximum open connections, and lifetime management via `database/sql` / `pgx`.
- **Safe Transaction Management:** Context-aware helper patterns to commit, rollback, and prevent dangling connections or leaks.
- **Core CRUD Operations:** Reusable query helpers for common PostgreSQL patterns (single record fetch, batched queries, inserts with `RETURNING`).
- **Context & Timeout Support:** Full context propagation to ensure queries fail fast under network partitions or heavy load.
- **Structured Error Handling:** Detailed logging and SQL error wrapping for simplified debugging and observability.
- **Environment-Driven Configuration:** Decoupled database credentials using environment variables.

---

## Project Structure

```text
.
├── config/             # Database connection configuration & environment loaders
├── db/                 # DB initialization, pool management, and ping checks
├── repository/         # Data access layer and CRUD implementation
├── migrations/         # SQL migration scripts (schema definition)
├── main.go             # Entrypoint demonstrating database connection and sample queries
├── go.mod              # Module dependencies
└── README.md
