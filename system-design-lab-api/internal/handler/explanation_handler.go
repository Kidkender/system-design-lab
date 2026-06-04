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

type ExplanationHandler struct {
	service *service.ExplanationService
}

func NewExplanationHandler(s *service.ExplanationService) *ExplanationHandler {
	return &ExplanationHandler{service: s}
}

// CreateExplanation godoc
// @Summary      Create an explanation
// @Tags         explanations
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateExplanationRequest  true  "Explanation payload"
// @Success      201   {object}  dto.ExplanationCreatedResponse
// @Failure      400   {string}  string
// @Failure      500   {string}  string
// @Router       /explanations [post]
func (h *ExplanationHandler) CreateExplanation(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateExplanationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		response.Error(w, err)
		return
	}

	resp, err := h.service.CreateExplanation(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusCreated, resp)
}

// UpdateExplanation godoc
// @Summary      Update an explanation
// @Tags         explanations
// @Accept       json
// @Produce      json
// @Param        id    path      string                       true  "Explanation UUID"
// @Param        body  body      dto.UpdateExplanationRequest true  "Explanation payload"
// @Success      200   {object}  dto.ExplanationCreatedResponse
// @Failure      400   {string}  string
// @Failure      500   {string}  string
// @Router       /explanations/{id} [put]
func (h *ExplanationHandler) UpdateExplanation(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(w, err)
		return
	}

	var req dto.UpdateExplanationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		response.Error(w, err)
		return
	}

	resp, err := h.service.UpdateExplanation(r.Context(), id, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, resp)
}

// DeleteExplanation godoc
// @Summary      Delete an explanation
// @Tags         explanations
// @Param        id   path      string  true  "Explanation UUID"
// @Success      204
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /explanations/{id} [delete]
func (h *ExplanationHandler) DeleteExplanation(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(w, err)
		return
	}

	if err := h.service.DeleteExplanation(r.Context(), id); err != nil {
		response.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ExplanationHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /explanations", h.CreateExplanation)
	mux.HandleFunc("PUT /explanations/{id}", h.UpdateExplanation)
	mux.HandleFunc("DELETE /explanations/{id}", h.DeleteExplanation)
}
