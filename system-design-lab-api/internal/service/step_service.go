package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type StepService struct {
	q *db.Queries
}

func NewStepService(q *db.Queries) *StepService {
	return &StepService{q: q}
}

func (s *StepService) CreateStep(ctx context.Context, req *dto.CreateStepRequest) (dto.StepResponse, error) {
	scenarioID, err := uuid.Parse(req.ScenarioID)
	if err != nil {
		return dto.StepResponse{}, err
	}

	_, err = s.q.GetScenario(ctx, scenarioID)
	if err != nil {
		slog.Error("Scenario not found", "error", err)
		return dto.StepResponse{}, err
	}

	step, err := s.q.CreateStep(ctx, db.CreateStepParams{
		Column1:    uuid.New(),
		Column2:    scenarioID,
		Question:   req.Question,
		Context:    pgtype.Text{String: *req.Context, Valid: req.Context != nil},
		OrderIndex: req.OrderIndex,
	})
	if err != nil {
		slog.Error("Failed to create step", "error", err)
		return dto.StepResponse{}, err
	}

	slog.Info("Step created successfully", "id", step.ID)
	return dto.StepResponse{
		ID:       step.ID.String(),
		Question: step.Question,
		Context:  &step.Context.String,
		Choices:  []dto.ChoiceResponse{},
	}, nil
}

func (s *StepService) GetStepsPaginated(
	ctx context.Context, page int32, limit int32,
) (*dto.PaginationResponse[dto.StepResponse], error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	total, err := s.q.CountSteps(ctx)
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	steps, err := s.q.GetStepsPaginated(ctx, db.GetStepsPaginatedParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	data := make([]dto.StepResponse, 0, len(steps))
	for _, step := range steps {
		data = append(data, dto.StepResponse{
			ID:       step.ID.String(),
			Question: step.Question,
			Context:  &step.Context.String,
			Choices:  nil,
		})
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return &dto.PaginationResponse[dto.StepResponse]{
		Data:       data,
		Total:      int32(total),
		Page:       page,
		Limit:      limit,
		TotalPages: int32(totalPages),
	}, nil

}
