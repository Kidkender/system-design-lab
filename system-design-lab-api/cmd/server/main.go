package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler"
	"github.com/kidkender/system-design-lab/internal/service"
)

func main() {
	conn, err := pgxpool.New(context.Background(), "postgres://root:root@localhost:5433/system_db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close()
	log.Println("connected to database")

	q := db.New(conn)
	service := service.NewScenarioService(q)
	handler := handler.NewScenarioHandler(service)

	http.HandleFunc("/scenarios/", handler.GetScenario)

	log.Println("server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
