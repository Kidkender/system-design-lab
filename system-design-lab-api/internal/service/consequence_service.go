package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type ConsequenceService struct {
	q *db.Queries
}

func NewConsequenceService(q *db.Queries) *ConsequenceService {
	return &ConsequenceService{q: q}
}

func (s *ConsequenceService) CreateConsequence(ctx context.Context, req *dto.CreateConsequenceRequest) (*dto.ConsequenceCreatedResponse, error) {
	choiceID, err := uuid.Parse(req.ChoiceID)
	if err != nil {
		return nil, err
	}

	row, err := s.q.CreateConsequence(ctx, db.CreateConsequenceParams{
		Column1: uuid.New(),
		Column2: choiceID,
		Type:    req.Type,
		Keys:    req.Keys,
		Value:   req.Value,
	})
	if err != nil {
		return nil, err
	}

	return &dto.ConsequenceCreatedResponse{
		ID:       row.ID.String(),
		ChoiceID: row.ChoiceID.String(),
		Type:     row.Type,
		Keys:     row.Keys,
		Value:    row.Value,
	}, nil
}

func (s *ConsequenceService) UpdateConsequence(ctx context.Context, id uuid.UUID, req *dto.UpdateConsequenceRequest) (*dto.ConsequenceCreatedResponse, error) {
	row, err := s.q.UpdateConsequence(ctx, db.UpdateConsequenceParams{
		Column1: id,
		Type:    req.Type,
		Keys:    req.Keys,
		Value:   req.Value,
	})
	if err != nil {
		return nil, err
	}

	return &dto.ConsequenceCreatedResponse{
		ID:       row.ID.String(),
		ChoiceID: row.ChoiceID.String(),
		Type:     row.Type,
		Keys:     row.Keys,
		Value:    row.Value,
	}, nil
}

func (s *ConsequenceService) DeleteConsequence(ctx context.Context, id uuid.UUID) error {
	return s.q.DeleteConsequence(ctx, id)
}
