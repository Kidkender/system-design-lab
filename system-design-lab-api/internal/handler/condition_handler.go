package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
	"github.com/kidkender/system-design-lab/internal/service"
	v "github.com/kidkender/system-design-lab/internal/validator"
)

type ConditionHandler struct {
	service *service.ConditionService
}

func NewConditionHandler(s *service.ConditionService) *ConditionHandler {
	return &ConditionHandler{service: s}
}

func (h *ConditionHandler) CreateCondition(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateConditionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.CreateCondition(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *ConditionHandler) DeleteCondition(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteCondition(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ConditionHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /conditions", h.CreateCondition)
	mux.HandleFunc("DELETE /conditions/{id}", h.DeleteCondition)
}
