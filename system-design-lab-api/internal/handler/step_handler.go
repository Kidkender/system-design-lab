package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kidkender/system-design-lab/internal/common/response"
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

// CreateStep godoc
// @Summary      Create a step
// @Tags         steps
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateStepRequest  true  "Step payload"
// @Success      201   {object}  dto.StepResponse
// @Failure      400   {string}  string
// @Failure      500   {string}  string
// @Router       /steps [post]
func (h *StepHandler) CreateStep(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateStepRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		response.Error(w, err)
		return
	}

	resp, err := h.service.CreateStep(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusCreated, resp)
}

// GetStepsPaginated godoc
// @Summary      List steps
// @Tags         steps
// @Produce      json
// @Param        page   query     int  false  "Page number"
// @Param        limit  query     int  false  "Items per page"
// @Success      200    {object}  dto.StepPaginationResponse
// @Failure      500    {string}  string
// @Router       /steps [get]
func (h *StepHandler) GetStepsPaginated(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	resp, err := h.service.GetStepsPaginated(r.Context(), int32(page), int32(limit))
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, resp)
}

func (h *StepHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /steps", h.GetStepsPaginated)
	mux.HandleFunc("POST /steps", h.CreateStep)
}
