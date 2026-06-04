package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/common/response"
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
		response.Error(w, err)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		response.Error(w, err)
		return
	}

	resp, err := h.service.CreateCondition(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusCreated, resp)
}

func (h *ConditionHandler) DeleteCondition(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(w, err)
		return
	}

	if err := h.service.DeleteCondition(r.Context(), id); err != nil {
		response.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ConditionHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /conditions", h.CreateCondition)
	mux.HandleFunc("DELETE /conditions/{id}", h.DeleteCondition)
}
