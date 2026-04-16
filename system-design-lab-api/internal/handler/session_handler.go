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
