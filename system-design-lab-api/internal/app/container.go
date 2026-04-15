package app

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler"
	"github.com/kidkender/system-design-lab/internal/service"
)

type Container struct {
	ScenarioHandler *handler.ScenarioHandler
	StepHandler     *handler.StepHandler
	ChoiceHandler   *handler.ChoiceHandler
	SessionHandler  *handler.SessionHandler
}

func NewContainer(conn *pgxpool.Pool) *Container {
	q := db.New(conn)

	scenarioService := service.NewScenarioService(q)
	stepService := service.NewStepService(q)
	choiceService := service.NewChoiceService(q)
	sessionService := service.NewSessionService(q)

	return &Container{
		ScenarioHandler: handler.NewScenarioHandler(scenarioService),
		StepHandler:     handler.NewStepHandler(stepService),
		ChoiceHandler:   handler.NewChoiceHandler(choiceService),
		SessionHandler:  handler.NewSessionHandler(sessionService),
	}
}

func (c *Container) RegisterRoutes(mux *http.ServeMux) {
	apiMux := http.NewServeMux()

	c.ScenarioHandler.RegisterRoutes(apiMux)
	c.StepHandler.RegisterRoutes(apiMux)
	c.ChoiceHandler.RegisterRoutes(apiMux)
	c.SessionHandler.RegisterRoutes(apiMux)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiMux))
}
