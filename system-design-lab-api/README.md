# system-design-lab-api

Backend API cho [System Design Lab](https://github.com/kidkender/system-design-lab) — nền tảng học system design thông qua quyết định và trade-off.

## Tech Stack

- **Go** 1.25
- **PostgreSQL** — primary database
- **pgx/v5** — PostgreSQL driver
- **sqlc** — type-safe SQL query generation

## Project Structure

```
system-design-lab-api/
├── cmd/server/         # Entry point (main.go)
├── db/                 # Raw SQL schema
│   └── schema.sql
├── internal/
│   ├── db/             # sqlc-generated models & queries
│   ├── handler/        # HTTP handlers
│   │   └── dto/        # Request/Response DTOs
│   └── service/        # Business logic
├── go.mod
├── go.sum
└── sqlc.yaml
```

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL running on port `5433`
- [sqlc](https://sqlc.dev/) (optional, for regenerating queries)

### Setup

1. Clone the repo

```bash
git clone https://github.com/kidkender/system-design-lab.git
cd system-design-lab/system-design-lab-api
```

2. Start PostgreSQL và tạo database

```bash
createdb system_db
psql -d system_db -f db/schema.sql
```

3. Run server

```bash
go run cmd/server/main.go
```

Server khởi động tại `http://localhost:8080`.

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/scenarios/:id` | Lấy thông tin scenario theo UUID |

### Example

```bash
GET /scenarios/550e8400-e29b-41d4-a716-446655440000
```

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Chat App",
  "description": "Design a real-time chat application",
  "difficulty": "medium",
  "steps": [...]
}
```

## Database Schema

Các bảng chính:

| Table | Mô tả |
|-------|-------|
| `scenarios` | Kịch bản học (Chat App, URL Shortener...) |
| `steps` | Các bước quyết định trong scenario |
| `choices` | Các lựa chọn cho mỗi step |
| `impacts` | Ảnh hưởng của choice lên metrics (latency, cost, scalability) |
| `explanations` | Giải thích 3 cấp độ (beginner / intermediate / advanced) |
| `consequences` | Hệ quả của lựa chọn (flag-based) |
| `users` | Người dùng |
| `user_sessions` | Phiên chơi của user |
| `user_choices` | Lịch sử lựa chọn của user |

## Development

### Regenerate sqlc queries

```bash
sqlc generate
```

### Build binary

```bash
go build -o system-design-lab ./cmd/server
```
