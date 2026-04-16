package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
	"github.com/kidkender/system-design-lab/internal/service"
	v "github.com/kidkender/system-design-lab/internal/validator"
)

type ConsequenceHandler struct {
	service *service.ConsequenceService
}

func NewConsequenceHandler(s *service.ConsequenceService) *ConsequenceHandler {
	return &ConsequenceHandler{service: s}
}

func (h *ConsequenceHandler) CreateConsequence(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateConsequenceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.CreateConsequence(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *ConsequenceHandler) UpdateConsequence(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req dto.UpdateConsequenceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.UpdateConsequence(r.Context(), id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *ConsequenceHandler) DeleteConsequence(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteConsequence(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ConsequenceHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /consequences", h.CreateConsequence)
	mux.HandleFunc("PUT /consequences/{id}", h.UpdateConsequence)
	mux.HandleFunc("DELETE /consequences/{id}", h.DeleteConsequence)
}
