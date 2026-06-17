package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/common/response"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
	"github.com/kidkender/system-design-lab/internal/service"
)

type ScenarioHandler struct {
	service        *service.ScenarioService
	sessionService *service.SessionService
}

func NewScenarioHandler(s *service.ScenarioService, sessionService *service.SessionService) *ScenarioHandler {
	return &ScenarioHandler{service: s, sessionService: sessionService}
}

// GetScenariosPaginated godoc
// @Summary      List scenarios
// @Tags         scenarios
// @Produce      json
// @Param        page   query     int  false  "Page number"
// @Param        limit  query     int  false  "Items per page"
// @Success      200    {object}  dto.ScenarioPaginationResponse
// @Failure      500    {string}  string
// @Router       /scenarios [get]
func (h *ScenarioHandler) GetScenariosPaginated(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 20
	}

	filter := dto.ScenarioFilter{
		Page:  page,
		Limit: limit,
	}
	if d := r.URL.Query().Get("difficulty"); d != "" {
		filter.Difficulty = &d
	}

	resp, err := h.service.GetScenariosPaginated(r.Context(), filter)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, resp)
}

// CreateScenario godoc
// @Summary      Create a scenario
// @Tags         scenarios
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateScenarioRequest  true  "Scenario payload"
// @Success      201
// @Failure      400   {string}  string
// @Failure      500   {string}  string
// @Router       /scenarios [post]
func (h *ScenarioHandler) CreateScenario(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateScenarioRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateScenario(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusCreated, map[string]string{"id": id.String()})
}

// GetScenario godoc
// @Summary      Get a scenario by ID
// @Tags         scenarios
// @Produce      json
// @Param        id   path      string  true  "Scenario UUID"
// @Success      200  {object}  dto.ScenarioResponse
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /scenarios/{id} [get]
func (h *ScenarioHandler) GetScenario(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(w, err)
		return
	}

	resp, err := h.service.GetScenario(r.Context(), id)

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, resp)
}

// GetLeaderboard godoc
// @Summary      Get leaderboard for a scenario
// @Tags         scenarios
// @Produce      json
// @Param        id     path      string  true   "Scenario UUID"
// @Param        top_n  query     int     false  "Number of entries (default 10)"
// @Success      200    {array}   dto.LeaderboardEntry
// @Failure      400    {string}  string
// @Router       /scenarios/{id}/leaderboard [get]
func (h *ScenarioHandler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	scenarioID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(w, err)
		return
	}

	topN := int32(10)
	if v := r.URL.Query().Get("top_n"); v != "" {
		n, err := strconv.Atoi(v)
		if err == nil && n > 0 {
			topN = int32(n)
		}
	}

	entries, err := h.sessionService.GetLeaderboard(r.Context(), scenarioID, topN)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, entries)
}

func (h *ScenarioHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /scenarios", h.GetScenariosPaginated)
	mux.HandleFunc("POST /scenarios", h.CreateScenario)
	mux.HandleFunc("GET /scenarios/{id}", h.GetScenario)
	mux.HandleFunc("GET /scenarios/{id}/leaderboard", h.GetLeaderboard)
}
