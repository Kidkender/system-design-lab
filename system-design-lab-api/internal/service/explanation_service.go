package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type ExplanationService struct {
	q *db.Queries
}

func NewExplanationService(q *db.Queries) *ExplanationService {
	return &ExplanationService{q: q}
}

func (s *ExplanationService) CreateExplanation(ctx context.Context, req *dto.CreateExplanationRequest) (*dto.ExplanationCreatedResponse, error) {
	choiceID, err := uuid.Parse(req.ChoiceID)
	if err != nil {
		return nil, err
	}

	exp, err := s.q.CreateExplanation(ctx, db.CreateExplanationParams{
		Column1: uuid.New(),
		Column2: choiceID,
		Level:   db.ExplanationLevel(req.Level),
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &dto.ExplanationCreatedResponse{
		ID:       exp.ID.String(),
		ChoiceID: exp.ChoiceID.String(),
		Level:    string(exp.Level),
		Content:  exp.Content,
	}, nil
}

func (s *ExplanationService) UpdateExplanation(ctx context.Context, id uuid.UUID, req *dto.UpdateExplanationRequest) (*dto.ExplanationCreatedResponse, error) {
	exp, err := s.q.UpdateExplanation(ctx, db.UpdateExplanationParams{
		Column1: id,
		Level:   db.ExplanationLevel(req.Level),
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &dto.ExplanationCreatedResponse{
		ID:       exp.ID.String(),
		ChoiceID: exp.ChoiceID.String(),
		Level:    string(exp.Level),
		Content:  exp.Content,
	}, nil
}

func (s *ExplanationService) DeleteExplanation(ctx context.Context, id uuid.UUID) error {
	return s.q.DeleteExplanation(ctx, id)
}
