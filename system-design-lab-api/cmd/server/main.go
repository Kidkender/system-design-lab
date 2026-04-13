package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler"
	"github.com/kidkender/system-design-lab/internal/service"
)

func main() {
	conn, err := pgxpool.New(context.Background(), "postgres://root:root@localhost:5433/system_db")
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
	}
	defer conn.Close()
	slog.Info("connected to database")

	mux := http.NewServeMux()

	q := db.New(conn)
	scenarioService := service.NewScenarioService(q)
	scenarioHandler := handler.NewScenarioHandler(scenarioService)
	stepService := service.NewStepService(q)
	stepHandler := handler.NewStepHandler(stepService)

	scenarioHandler.RegisterRoutes(mux)
	stepHandler.RegisterRoutes(mux)

	slog.Info("server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("server error", "error", err)
	}
}
