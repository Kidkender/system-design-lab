package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kidkender/system-design-lab/internal/handler/dto"
	"github.com/kidkender/system-design-lab/internal/service"

	v "github.com/kidkender/system-design-lab/internal/validator"
)

type StepHandler struct {
	service *service.StepService
}

func NewStepHandler(service *service.StepService) *StepHandler {
	return &StepHandler{service: service}
}

func (h *StepHandler) CreateStep(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateStepRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.CreateStep(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *StepHandler) GetStepsPaginated(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	resp, err := h.service.GetStepsPaginated(r.Context(), int32(page), int32(limit))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *StepHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /steps", h.GetStepsPaginated)
	mux.HandleFunc("POST /steps", h.CreateStep)
}
