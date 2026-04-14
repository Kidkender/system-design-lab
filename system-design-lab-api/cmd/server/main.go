package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kidkender/system-design-lab/internal/app"
)

func main() {
	conn, err := pgxpool.New(context.Background(), "postgres://root:root@localhost:5433/system_db")
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
	}
	defer conn.Close()
	slog.Info("connected to database")

	mux := http.NewServeMux()

	container := app.NewContainer(conn)
	container.RegisterRoutes(mux)

	slog.Info("server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("server error", "error", err)
	}
}
