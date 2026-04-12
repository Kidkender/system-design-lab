package service

import (
	"context"

	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type StepService struct {
	q *db.Queries
}

func NewStepService(q *db.Queries) *StepService {
	return &StepService{q: q}
}

func (s *StepService) CreateStep(ctx context.Context, req dto.)