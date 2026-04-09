package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/service"
)

type ScenarioHandler struct {
	service *service.ScenarioService
}

func NewScenarioHandler(s *service.ScenarioService) *ScenarioHandler {
	return &ScenarioHandler{service: s}
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
