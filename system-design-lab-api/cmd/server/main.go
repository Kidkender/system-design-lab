package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/kidkender/system-design-lab/docs"
	"github.com/kidkender/system-design-lab/internal/app"
	"github.com/kidkender/system-design-lab/internal/middleware"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://root:root@localhost:5433/system_db"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer conn.Close()
	slog.Info("connected to database")

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		if err := conn.Ping(r.Context()); err != nil {
			http.Error(w, "db unreachable", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	container := app.NewContainer(conn)
	container.RegisterRoutes(mux)

	corsMux := middleware.EnableCORS(mux)
	loggedMux := middleware.LoggingMiddleware(corsMux)
	rateLimitedMux := middleware.RateLimitMiddleware(loggedMux)

	slog.Info("server running", "port", port)
	if err := http.ListenAndServe(":"+port, rateLimitedMux); err != nil {
		slog.Error("server error", "error", err)
	}
}
