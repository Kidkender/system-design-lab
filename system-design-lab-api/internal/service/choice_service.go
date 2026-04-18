package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type choiceQuerier interface {
	CreateChoice(ctx context.Context, arg db.CreateChoiceParams) (db.Choice, error)
}

type ChoiceService struct {
	q choiceQuerier
}

func NewChoiceService(q choiceQuerier) *ChoiceService {
	return &ChoiceService{q: q}
}

func (s *ChoiceService) CreateChoice(ctx context.Context, req *dto.CreateChoiceRequest) (*dto.ChoiceCreatedResponse, error) {
	stepID, err := uuid.Parse(req.StepID)
	if err != nil {
		return nil, err
	}

	nextStepID := uuid.Nil
	if req.NextStepID != nil {
		nextStepID, err = uuid.Parse(*req.NextStepID)
		if err != nil {
			return nil, err
		}
	}

	choice, err := s.q.CreateChoice(ctx, db.CreateChoiceParams{
		ID:         uuid.New(),
		StepID:     stepID,
		Label:      req.Label,
		NextStepID: nextStepID,
		IsCorrect:  req.IsCorrect,
	})
	if err != nil {
		return nil, err
	}

	resp := &dto.ChoiceCreatedResponse{
		ID:        choice.ID.String(),
		StepID:    choice.StepID.String(),
		Label:     choice.Label,
		IsCorrect: choice.IsCorrect,
	}
	if choice.NextStepID != uuid.Nil {
		s := choice.NextStepID.String()
		resp.NextStepID = &s
	}

	return resp, nil
}
