package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type SessionService struct {
	q *db.Queries
}

func NewSessionService(q *db.Queries) *SessionService {
	return &SessionService{q: q}
}

func (s *SessionService) StartSession(ctx context.Context, req *dto.StartSessionRequest) (*dto.SessionResponse, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}

	scenarioID, err := uuid.Parse(req.ScenarioID)
	if err != nil {
		return nil, err
	}

	scenario, err := s.q.GetScenario(ctx, scenarioID)
	if err != nil {
		return nil, fmt.Errorf("scenario not found: %w", err)
	}

	if scenario.StartStepID == uuid.Nil {
		return nil, errors.New("scenario has no start step")
	}

	metricsJSON, _ := json.Marshal(map[string]float64{
		"latency":     0,
		"cost":        0,
		"scalability": 0,
	})
	flagsJSON, _ := json.Marshal(map[string]bool{})

	session, err := s.q.CreateUserSession(ctx, db.CreateUserSessionParams{
		ID:            uuid.New(),
		UserID:        userID,
		ScenarioID:    scenarioID,
		CurrentStepID: scenario.StartStepID,
		Metrics:       metricsJSON,
		Flags:         flagsJSON,
	})
	if err != nil {
		return nil, err
	}

	stepResp, err := s.buildStepResponse(ctx, scenario.StartStepID)
	if err != nil {
		return nil, err
	}

	var metrics map[string]float64
	var flags map[string]bool
	json.Unmarshal(session.Metrics, &metrics)
	json.Unmarshal(session.Flags, &flags)

	return &dto.SessionResponse{
		ID:          session.ID.String(),
		ScenarioID:  session.ScenarioID.String(),
		CurrentStep: *stepResp,
		Metrics:     metrics,
		Flags:       flags,
		Status:      string(session.Status),
	}, nil
}

func (s *SessionService) SubmitChoice(ctx context.Context, sessionID uuid.UUID, req *dto.SubmitChoiceRequest) (*dto.SubmitChoiceResponse, error) {
	session, err := s.q.GetUserSession(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	if session.Status != db.SessionStatusInProgress {
		return nil, errors.New("session is already completed")
	}

	choiceID, err := uuid.Parse(req.ChoiceID)
	if err != nil {
		return nil, err
	}

	choice, err := s.q.GetChoice(ctx, choiceID)
	if err != nil {
		return nil, fmt.Errorf("choice not found: %w", err)
	}

	if choice.StepID != session.CurrentStepID {
		return nil, errors.New("choice does not belong to current step")
	}

	// Record user choice
	_, err = s.q.CreateUserChoice(ctx, db.CreateUserChoiceParams{
		ID:        uuid.New(),
		SessionID: session.ID,
		StepID:    session.CurrentStepID,
		ChoiceID:  choiceID,
	})
	if err != nil {
		return nil, err
	}

	// Apply impacts to metrics
	var metrics map[string]float64
	json.Unmarshal(session.Metrics, &metrics)

	impacts, _ := s.q.GetImpactsByChoiceIDs(ctx, []uuid.UUID{choiceID})
	for _, impact := range impacts {
		metric := string(impact.Metric)
		switch impact.Type {
		case db.ImpactTypeAdd:
			metrics[metric] += impact.Value
		case db.ImpactTypeMultiply:
			metrics[metric] *= impact.Value
		case db.ImpactTypeSet:
			metrics[metric] = impact.Value
		}
	}

	metricsJSON, _ := json.Marshal(metrics)

	// Determine next step and status
	nextStepID := choice.NextStepID
	status := db.SessionStatusInProgress
	isCompleted := false

	if nextStepID == uuid.Nil {
		status = db.SessionStatusCompleted
		isCompleted = true
	}

	_, err = s.q.UpdateUserSession(ctx, db.UpdateUserSessionParams{
		ID:            session.ID,
		CurrentStepID: nextStepID,
		Metrics:       metricsJSON,
		Flags:         session.Flags,
		Status:        status,
	})
	if err != nil {
		return nil, err
	}

	resp := &dto.SubmitChoiceResponse{
		IsCorrect:   choice.IsCorrect,
		Metrics:     metrics,
		IsCompleted: isCompleted,
	}

	if choice.IsCorrect {
		resp.Feedback = "✅ Good choice!"
	} else {
		resp.Feedback = "❌ Not the best choice for this scenario."
	}

	if !isCompleted {
		nextStep, err := s.buildStepResponse(ctx, nextStepID)
		if err != nil {
			return nil, err
		}
		resp.NextStep = nextStep
	}

	return resp, nil
}

func (s *SessionService) buildStepResponse(ctx context.Context, stepID uuid.UUID) (*dto.StepResponse, error) {
	step, err := s.q.GetStep(ctx, stepID)
	if err != nil {
		return nil, err
	}

	choices, err := s.q.GetChoicesByStepIDs(ctx, []uuid.UUID{stepID})
	if err != nil {
		return nil, err
	}

	choiceResps := make([]dto.ChoiceResponse, 0, len(choices))
	for _, c := range choices {
		nextStepID := c.NextStepID.String()
		ch := dto.ChoiceResponse{
			ID:    c.ID.String(),
			Label: c.Label,
		}
		if c.NextStepID != uuid.Nil {
			ch.NextStepID = &nextStepID
		}
		choiceResps = append(choiceResps, ch)
	}

	ctx2 := step.Context.String
	return &dto.StepResponse{
		ID:       step.ID.String(),
		Question: step.Question,
		Context:  &ctx2,
		Choices:  choiceResps,
	}, nil
}
