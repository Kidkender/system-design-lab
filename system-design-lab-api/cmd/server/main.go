package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/kidkender/system-design-lab/docs"
	"github.com/kidkender/system-design-lab/internal/app"
	"github.com/kidkender/system-design-lab/internal/middleware"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	conn, err := pgxpool.New(context.Background(), "postgres://root:root@localhost:5433/system_db")
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
	}
	defer conn.Close()
	slog.Info("connected to database")

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	container := app.NewContainer(conn)
	container.RegisterRoutes(mux)

	corsMux := middleware.EnableCORS(mux)
	loggedMux := middleware.LoggingMiddleware(corsMux)
	slog.Info("server running on :8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		slog.Error("server error", "error", err)
	}
}
