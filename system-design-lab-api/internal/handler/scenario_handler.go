package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
	"github.com/kidkender/system-design-lab/internal/service"
)

type ScenarioHandler struct {
	service *service.ScenarioService
}

func NewScenarioHandler(s *service.ScenarioService) *ScenarioHandler {
	return &ScenarioHandler{service: s}
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
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 20
	}

	resp, err := h.service.GetScenariosPaginated(r.Context(), int32(page), int32(limit))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
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

	h.service.CreateScenario(r.Context(), &req)

	w.WriteHeader(http.StatusCreated)
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
	idStr := strings.TrimPrefix(r.URL.Path, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	resp, err := h.service.GetScenario(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *ScenarioHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /scenarios", h.GetScenariosPaginated)
	mux.HandleFunc("POST /scenarios", h.CreateScenario)
	mux.HandleFunc("GET /scenarios/{id}", h.GetScenario)
}
