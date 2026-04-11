include .env

migrate-up:
	migrate -path ./db/migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path ./db/migrations -database "$(DATABASE_URL)" down

migrate-force:
	migrate -path ./db/migrations -database "$(DATABASE_URL)" force $(VERSION)

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(NAME)

sqlc:
	sqlc generate

run:
	go run ./cmd/server/main.go
