# system-design-lab-api

Go backend cho [System Design Lab](../README.md).

## Yêu cầu

- Go 1.21+
- PostgreSQL (chạy trên port `5433`)
- [sqlc](https://sqlc.dev/) — chỉ cần khi thay đổi SQL schema/query

## Chạy nhanh

```bash
# 1. Tạo database
createdb system_db
psql -d system_db -f db/schema.sql

# 2. (Nếu DB đang có từ trước, chạy migration để thêm tính năng mới)
psql -d system_db -f db/migrations/001_add_features.sql

# 3. Chạy server
go run cmd/server/main.go
```

Server chạy tại `http://localhost:8080`.

## Biến môi trường

| Biến | Mặc định | Mô tả |
|------|----------|-------|
| `DATABASE_URL` | `postgres://root:root@localhost:5433/system_db` | PostgreSQL DSN |
| `PORT` | `8080` | Port server lắng nghe |

## API

| Method | Path | Mô tả |
|--------|------|-------|
| GET | `/healthz` | Health check |
| GET | `/api/v1/scenarios` | Danh sách scenario (có phân trang, filter difficulty) |
| GET | `/api/v1/scenarios/:id` | Chi tiết scenario |
| POST | `/api/v1/users` | Tạo user |
| POST | `/api/v1/sessions` | Bắt đầu session mới |
| GET | `/api/v1/sessions/:id` | Lấy trạng thái session |
| POST | `/api/v1/sessions/:id/submit` | Submit một lựa chọn |
| GET | `/api/v1/sessions/:id/summary` | Tổng kết session |
| POST | `/api/v1/sessions/:id/abandon` | Bỏ dở session |
| POST | `/api/v1/sessions/:id/restart` | Chơi lại từ đầu |
| GET | `/api/v1/scenarios/:id/leaderboard` | Leaderboard của scenario |
| GET | `/api/v1/users/:id/sessions` | Lịch sử session của user |
| GET | `/api/v1/users/:id/progress` | Tiến độ user qua các scenario |

## Lệnh thường dùng

```bash
# Chạy server
go run cmd/server/main.go

# Build binary
go build -o system-design-lab ./cmd/server

# Chạy tests
go test ./...

# Regenerate sqlc (sau khi sửa db/schema.sql hoặc db/query/*.sql)
sqlc generate
```

## Swagger

Docs tại `http://localhost:8080/swagger/index.html` khi server đang chạy.
