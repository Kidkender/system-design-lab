package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kidkender/system-design-lab/internal/common/response"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
	"github.com/kidkender/system-design-lab/internal/service"
	v "github.com/kidkender/system-design-lab/internal/validator"
)

type ChoiceHandler struct {
	service *service.ChoiceService
}

func NewChoiceHandler(s *service.ChoiceService) *ChoiceHandler {
	return &ChoiceHandler{service: s}
}

// CreateChoice godoc
// @Summary      Create a choice
// @Tags         choices
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateChoiceRequest  true  "Choice payload"
// @Success      201   {object}  dto.ChoiceCreatedResponse
// @Failure      400   {string}  string
// @Failure      500   {string}  string
// @Router       /choices [post]
func (h *ChoiceHandler) CreateChoice(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateChoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		response.Error(w, err)
		return
	}

	resp, err := h.service.CreateChoice(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusCreated, resp)
}

func (h *ChoiceHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /choices", h.CreateChoice)
}
