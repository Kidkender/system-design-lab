package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type ImpactService struct {
	q *db.Queries
}

func NewImpactService(q *db.Queries) *ImpactService {
	return &ImpactService{q: q}
}

func (s *ImpactService) CreateImpact(ctx context.Context, req *dto.CreateImpactRequest) (*dto.ImpactCreatedResponse, error) {
	choiceID, err := uuid.Parse(req.ChoiceID)
	if err != nil {
		return nil, err
	}

	impact, err := s.q.CreateImpact(ctx, db.CreateImpactParams{
		Column1: uuid.New(),
		Column2: choiceID,
		Metric:  db.ImpactMetric(req.Metric),
		Type:    db.ImpactType(req.Type),
		Value:   req.Value,
	})
	if err != nil {
		return nil, err
	}

	return &dto.ImpactCreatedResponse{
		ID:       impact.ID.String(),
		ChoiceID: impact.ChoiceID.String(),
		Metric:   string(impact.Metric),
		Type:     string(impact.Type),
		Value:    impact.Value,
	}, nil
}

func (s *ImpactService) UpdateImpact(ctx context.Context, id uuid.UUID, req *dto.UpdateImpactRequest) (*dto.ImpactCreatedResponse, error) {
	impact, err := s.q.UpdateImpact(ctx, db.UpdateImpactParams{
		Column1: id,
		Metric:  db.ImpactMetric(req.Metric),
		Type:    db.ImpactType(req.Type),
		Value:   req.Value,
	})
	if err != nil {
		return nil, err
	}

	return &dto.ImpactCreatedResponse{
		ID:       impact.ID.String(),
		ChoiceID: impact.ChoiceID.String(),
		Metric:   string(impact.Metric),
		Type:     string(impact.Type),
		Value:    impact.Value,
	}, nil
}

func (s *ImpactService) DeleteImpact(ctx context.Context, id uuid.UUID) error {
	return s.q.DeleteImpact(ctx, id)
}
