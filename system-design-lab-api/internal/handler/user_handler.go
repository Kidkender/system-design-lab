package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/common/response"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
	"github.com/kidkender/system-design-lab/internal/service"
)

type UserHandler struct {
	service        *service.UserService
	sessionService *service.SessionService
}

func NewUserHandler(s *service.UserService, sessionService *service.SessionService) *UserHandler {
	return &UserHandler{service: s, sessionService: sessionService}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}

	user, err := h.service.CreateUser(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusCreated, user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(w, err)
		return
	}

	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, user)
}

func (h *UserHandler) GetUserProgress(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(w, err)
		return
	}

	items, err := h.sessionService.GetUserProgress(r.Context(), userID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, items)
}

func (h *UserHandler) ListUserSessions(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(w, err)
		return
	}

	items, err := h.sessionService.ListSessionsByUser(r.Context(), userID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, items)
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", h.CreateUser)
	mux.HandleFunc("GET /users/{id}", h.GetUser)
	mux.HandleFunc("GET /users/{id}/sessions", h.ListUserSessions)
	mux.HandleFunc("GET /users/{id}/progress", h.GetUserProgress)
}
