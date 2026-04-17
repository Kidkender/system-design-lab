package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
	"github.com/kidkender/system-design-lab/internal/service"
	v "github.com/kidkender/system-design-lab/internal/validator"
)

type SessionHandler struct {
	service *service.SessionService
}

func NewSessionHandler(s *service.SessionService) *SessionHandler {
	return &SessionHandler{service: s}
}

// StartSession godoc
// @Summary      Start a new session
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Param        body  body      dto.StartSessionRequest  true  "Session payload"
// @Success      201   {object}  dto.SessionResponse
// @Failure      400   {string}  string
// @Failure      500   {string}  string
// @Router       /sessions [post]
func (h *SessionHandler) StartSession(w http.ResponseWriter, r *http.Request) {
	var req dto.StartSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.StartSession(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// SubmitChoice godoc
// @Summary      Submit a choice for a session
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Param        id    path      string                  true  "Session UUID"
// @Param        body  body      dto.SubmitChoiceRequest true  "Choice payload"
// @Success      200   {object}  dto.SubmitChoiceResponse
// @Failure      400   {string}  string
// @Failure      500   {string}  string
// @Router       /sessions/{id}/submit [post]
func (h *SessionHandler) SubmitChoice(w http.ResponseWriter, r *http.Request) {
	sessionIDStr := r.PathValue("id")
	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		http.Error(w, "invalid session id", http.StatusBadRequest)
		return
	}

	var req dto.SubmitChoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := v.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.SubmitChoice(r.Context(), sessionID, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GetSession godoc
// @Summary      Get session by ID
// @Tags         sessions
// @Produce      json
// @Param        id   path      string  true  "Session UUID"
// @Success      200  {object}  dto.SessionResponse
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /sessions/{id} [get]
func (h *SessionHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	sessionID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid session id", http.StatusBadRequest)
		return
	}

	resp, err := h.service.GetSession(r.Context(), sessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GetSessionSummary godoc
// @Summary      Get session summary
// @Tags         sessions
// @Produce      json
// @Param        id   path      string  true  "Session UUID"
// @Success      200  {object}  dto.SessionSummaryResponse
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /sessions/{id}/summary [get]
func (h *SessionHandler) GetSessionSummary(w http.ResponseWriter, r *http.Request) {
	sessionID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid session id", http.StatusBadRequest)
		return
	}

	resp, err := h.service.GetSessionSummary(r.Context(), sessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *SessionHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /sessions", h.StartSession)
	mux.HandleFunc("GET /sessions/{id}", h.GetSession)
	mux.HandleFunc("POST /sessions/{id}/submit", h.SubmitChoice)
	mux.HandleFunc("GET /sessions/{id}/summary", h.GetSessionSummary)
}
