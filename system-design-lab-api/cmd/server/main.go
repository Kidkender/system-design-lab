package main

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler"
	"github.com/kidkender/system-design-lab/internal/service"
)

func main() {
	conn, _ := pgxpool.New(context.Background(), "postgres://root:root@localhost:5433/system_db")
	q := db.New(conn)

	service := service.NewScenarioService(q)
	handler := handler.NewScenarioHandler(service)
	http.HandleFunc("/scenarios/", handler.GetScenario)
	http.ListenAndServe(":8080", nil)

}
