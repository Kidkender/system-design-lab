package handler

import (
	"encoding/json"
	"fmt"
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

func (h *ScenarioHandler) GetScenariosPaginated(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	fmt.Printf("page: %s, limit: %s\n", pageStr, limitStr)
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	resp, err := h.service.GetScenariosPaginated(r.Context(), int32(page), int32(limit))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *ScenarioHandler) CreateScenario(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateScenarioRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	h.service.CreateScenario(r.Context(), &req)

	w.WriteHeader(http.StatusCreated)
}

func (h *ScenarioHandler) GetScenario(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/scenarios/")
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
