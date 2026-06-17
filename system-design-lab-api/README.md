# system-design-lab-api

Go backend for [System Design Lab](../README.md).

## Requirements

- Go 1.21+
- PostgreSQL running on port `5433`
- [sqlc](https://sqlc.dev/) — only needed when changing SQL schema or queries

## Quick Start

```bash
# 1. Create the database
createdb system_db
psql -d system_db -f db/schema.sql

# 2. If upgrading an existing database, run the migration
psql -d system_db -f db/migrations/001_add_features.sql

# 3. Start the server
go run cmd/server/main.go
```

Server runs at `http://localhost:8080`.

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DATABASE_URL` | `postgres://root:root@localhost:5433/system_db` | PostgreSQL DSN |
| `PORT` | `8080` | Port the server listens on |

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/healthz` | Health check |
| GET | `/api/v1/scenarios` | List scenarios (paginated, filterable by difficulty) |
| GET | `/api/v1/scenarios/:id` | Get scenario detail |
| POST | `/api/v1/users` | Create a user |
| POST | `/api/v1/sessions` | Start a new session |
| GET | `/api/v1/sessions/:id` | Get session state |
| POST | `/api/v1/sessions/:id/submit` | Submit a choice |
| GET | `/api/v1/sessions/:id/summary` | Get session summary |
| POST | `/api/v1/sessions/:id/abandon` | Abandon a session |
| POST | `/api/v1/sessions/:id/restart` | Restart a session |
| GET | `/api/v1/scenarios/:id/leaderboard` | Scenario leaderboard |
| GET | `/api/v1/users/:id/sessions` | User session history |
| GET | `/api/v1/users/:id/progress` | User progress across scenarios |

## Commands

```bash
# Run the server
go run cmd/server/main.go

# Build binary
go build -o system-design-lab ./cmd/server

# Run tests
go test ./...

# Regenerate sqlc (after editing db/schema.sql or db/query/*.sql)
sqlc generate
```

## Swagger

API docs available at `http://localhost:8080/swagger/index.html` when the server is running.
