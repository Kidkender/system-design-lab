package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type SessionService struct {
	q    *db.Queries
	pool *pgxpool.Pool
}

func NewSessionService(q *db.Queries, pool *pgxpool.Pool) *SessionService {
	return &SessionService{q: q, pool: pool}
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

	// Record user choice (will be committed inside the transaction below)
	newChoiceID := uuid.New()

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

	// Apply consequences to flags
	var flags map[string]bool
	json.Unmarshal(session.Flags, &flags)
	if flags == nil {
		flags = map[string]bool{}
	}

	consequences, _ := s.q.GetConsequencesByChoiceIDs(ctx, []uuid.UUID{choiceID})
	for _, c := range consequences {
		flags[c.Keys] = c.Value
	}

	flagsJSON, _ := json.Marshal(flags)

	// Fetch explanations for the chosen option
	explanationRows, _ := s.q.GetExplanationsByChoiceIDs(ctx, []uuid.UUID{choiceID})
	explanations := make(map[string]string, len(explanationRows))
	for _, e := range explanationRows {
		explanations[string(e.Level)] = e.Content
	}

	// Collect prior choice IDs for condition evaluation (includes current choice)
	priorChoices, _ := s.q.GetUserChoicesBySession(ctx, session.ID)
	priorChoiceIDs := make(map[uuid.UUID]bool, len(priorChoices)+1)
	for _, pc := range priorChoices {
		priorChoiceIDs[pc.ChoiceID] = true
	}
	priorChoiceIDs[choiceID] = true

	// Determine next step and status
	nextStepID := choice.NextStepID
	status := db.SessionStatusInProgress
	isCompleted := false

	if nextStepID == uuid.Nil {
		status = db.SessionStatusCompleted
		isCompleted = true
	} else {
		ok, err := s.evaluateConditions(ctx, nextStepID, metrics, flags, priorChoiceIDs)
		if err != nil {
			return nil, fmt.Errorf("evaluating conditions: %w", err)
		}
		if !ok {
			// Next step conditions not met — end the session
			status = db.SessionStatusCompleted
			isCompleted = true
			nextStepID = uuid.Nil
		}
	}

	// Persist choice + session update atomically
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx) //nolint:errcheck

	qtx := s.q.WithTx(tx)

	_, err = qtx.CreateUserChoice(ctx, db.CreateUserChoiceParams{
		ID:        newChoiceID,
		SessionID: session.ID,
		StepID:    session.CurrentStepID,
		ChoiceID:  choiceID,
	})
	if err != nil {
		return nil, fmt.Errorf("record user choice: %w", err)
	}

	_, err = qtx.UpdateUserSession(ctx, db.UpdateUserSessionParams{
		ID:            session.ID,
		CurrentStepID: nextStepID,
		Metrics:       metricsJSON,
		Flags:         flagsJSON,
		Status:        status,
	})
	if err != nil {
		return nil, fmt.Errorf("update session: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	// Build feedback from beginner-level explanation; fall back to generic text
	feedback := ""
	if text, ok := explanations["beginner"]; ok && text != "" {
		if choice.IsCorrect {
			feedback = "✅ " + text
		} else {
			feedback = "❌ " + text
		}
	} else {
		if choice.IsCorrect {
			feedback = "✅ Good choice!"
		} else {
			feedback = "❌ Not the best choice for this scenario."
		}
	}

	resp := &dto.SubmitChoiceResponse{
		IsCorrect:    choice.IsCorrect,
		Feedback:     feedback,
		Explanations: explanations,
		Metrics:      metrics,
		IsCompleted:  isCompleted,
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

func (s *SessionService) GetSession(ctx context.Context, sessionID uuid.UUID) (*dto.SessionResponse, error) {
	session, err := s.q.GetUserSession(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	var metrics map[string]float64
	var flags map[string]bool
	json.Unmarshal(session.Metrics, &metrics)
	json.Unmarshal(session.Flags, &flags)

	resp := &dto.SessionResponse{
		ID:         session.ID.String(),
		ScenarioID: session.ScenarioID.String(),
		Metrics:    metrics,
		Flags:      flags,
		Status:     string(session.Status),
	}

	if session.CurrentStepID != uuid.Nil {
		stepResp, err := s.buildStepResponse(ctx, session.CurrentStepID)
		if err != nil {
			return nil, err
		}
		resp.CurrentStep = *stepResp
	}

	return resp, nil
}

func (s *SessionService) GetSessionSummary(ctx context.Context, sessionID uuid.UUID) (*dto.SessionSummaryResponse, error) {
	session, err := s.q.GetUserSession(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	var metrics map[string]float64
	var flags map[string]bool
	json.Unmarshal(session.Metrics, &metrics)
	json.Unmarshal(session.Flags, &flags)

	userChoices, err := s.q.GetUserChoicesBySession(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	// Batch-fetch choices and explanations
	choiceIDs := make([]uuid.UUID, 0, len(userChoices))
	for _, uc := range userChoices {
		choiceIDs = append(choiceIDs, uc.ChoiceID)
	}

	choiceRows, _ := s.q.GetChoicesByStepIDs(ctx, func() []uuid.UUID {
		ids := make([]uuid.UUID, 0, len(userChoices))
		for _, uc := range userChoices {
			ids = append(ids, uc.StepID)
		}
		return ids
	}())
	choiceMap := make(map[uuid.UUID]db.GetChoicesByStepIDsRow, len(choiceRows))
	for _, c := range choiceRows {
		choiceMap[c.ID] = c
	}

	explanationRows, _ := s.q.GetExplanationsByChoiceIDs(ctx, choiceIDs)
	explainMap := make(map[uuid.UUID]map[string]string)
	for _, e := range explanationRows {
		if explainMap[e.ChoiceID] == nil {
			explainMap[e.ChoiceID] = map[string]string{}
		}
		explainMap[e.ChoiceID][string(e.Level)] = e.Content
	}

	seen := make(map[uuid.UUID]bool)
	stepIDs := make(uuid.UUIDs, 0, len(userChoices))
	for _, uc := range userChoices {
		if !seen[uc.StepID] {
			seen[uc.StepID] = true
			stepIDs = append(stepIDs, uc.StepID)
		}
	}
	steps, err := s.q.GetStepsByIDs(ctx, stepIDs)
	if err != nil {
		return nil, err
	}
	stepQuestionMap := make(map[uuid.UUID]string, len(steps))
	for _, step := range steps {
		stepQuestionMap[step.ID] = step.Question
	}

	summaries := make([]dto.UserChoiceSummary, 0, len(userChoices))
	for _, uc := range userChoices {
		ch := choiceMap[uc.ChoiceID]
		summaries = append(summaries, dto.UserChoiceSummary{
			StepID:       uc.StepID.String(),
			Question:     stepQuestionMap[uc.StepID],
			ChoiceID:     uc.ChoiceID.String(),
			Label:        ch.Label,
			IsCorrect:    ch.IsCorrect,
			Explanations: explainMap[uc.ChoiceID],
		})
	}

	return &dto.SessionSummaryResponse{
		ID:         session.ID.String(),
		ScenarioID: session.ScenarioID.String(),
		Status:     string(session.Status),
		Metrics:    metrics,
		Flags:      flags,
		Choices:    summaries,
		CreatedAt:  session.CreatedAt.Time.Format(time.RFC3339),
	}, nil
}

// evaluateConditions checks whether all conditions on a step are satisfied.
// Returns true if there are no conditions or all conditions pass (AND logic).
func (s *SessionService) evaluateConditions(
	ctx context.Context,
	stepID uuid.UUID,
	metrics map[string]float64,
	flags map[string]bool,
	priorChoiceIDs map[uuid.UUID]bool,
) (bool, error) {
	conditions, err := s.q.GetConditionsByStep(ctx, stepID)
	if err != nil {
		return false, err
	}
	if len(conditions) == 0 {
		return true, nil
	}

	for _, c := range conditions {
		switch c.Type {
		case db.ConditionTypeMetric:
			if !c.Metric.Valid || !c.Operator.Valid || !c.Value.Valid {
				continue
			}
			actual := metrics[string(c.Metric.ImpactMetric)]
			threshold := c.Value.Float64
			var pass bool
			switch c.Operator.String {
			case ">":
				pass = actual > threshold
			case "<":
				pass = actual < threshold
			case ">=":
				pass = actual >= threshold
			case "<=":
				pass = actual <= threshold
			case "==":
				pass = actual == threshold
			case "!=":
				pass = actual != threshold
			default:
				continue
			}
			if !pass {
				return false, nil
			}

		case db.ConditionTypeFlag:
			if !c.FloatKey.Valid {
				continue
			}
			if !flags[c.FloatKey.String] {
				return false, nil
			}

		case db.ConditionTypeChoice:
			if c.ChoiceID == uuid.Nil {
				continue
			}
			if !priorChoiceIDs[c.ChoiceID] {
				return false, nil
			}
		}
	}
	return true, nil
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

// func (s *SessionService) GetSessionByUserID(ctx context.Context, userID uuid.UUID) (*dto.SessionResponse, error) {
// 	session, err := s.q.GetSession(ctx, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// }
