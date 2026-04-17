package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
	"github.com/kidkender/system-design-lab/internal/service"
	v "github.com/kidkender/system-design-lab/internal/validator"
)

type ImpactHandler struct {
	service *service.ImpactService
}

func NewImpactHandler(s *service.ImpactService) *ImpactHandler {
	return &ImpactHandler{service: s}
}

// CreateImpact godoc
// @Summary      Create an impact
// @Tags         impacts
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateImpactRequest  true  "Impact payload"
// @Success      201   {object}  dto.ImpactCreatedResponse
// @Failure      400   {string}  string
// @Failure      500   {string}  string
// @Router       /impacts [post]
func (h *ImpactHandler) CreateImpact(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateImpactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.CreateImpact(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// UpdateImpact godoc
// @Summary      Update an impact
// @Tags         impacts
// @Accept       json
// @Produce      json
// @Param        id    path      string                  true  "Impact UUID"
// @Param        body  body      dto.UpdateImpactRequest true  "Impact payload"
// @Success      200   {object}  dto.ImpactCreatedResponse
// @Failure      400   {string}  string
// @Failure      500   {string}  string
// @Router       /impacts/{id} [put]
func (h *ImpactHandler) UpdateImpact(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req dto.UpdateImpactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.UpdateImpact(r.Context(), id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DeleteImpact godoc
// @Summary      Delete an impact
// @Tags         impacts
// @Param        id   path      string  true  "Impact UUID"
// @Success      204
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /impacts/{id} [delete]
func (h *ImpactHandler) DeleteImpact(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteImpact(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ImpactHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /impacts", h.CreateImpact)
	mux.HandleFunc("PUT /impacts/{id}", h.UpdateImpact)
	mux.HandleFunc("DELETE /impacts/{id}", h.DeleteImpact)
}
