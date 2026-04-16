package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type ConditionService struct {
	q *db.Queries
}

func NewConditionService(q *db.Queries) *ConditionService {
	return &ConditionService{q: q}
}

func (s *ConditionService) CreateCondition(ctx context.Context, req *dto.CreateConditionRequest) (*dto.ConditionCreatedResponse, error) {
	stepID, err := uuid.Parse(req.StepID)
	if err != nil {
		return nil, err
	}

	var metric db.ImpactMetric
	if req.Metric != nil {
		metric = db.ImpactMetric(*req.Metric)
	}

	var operator string
	if req.Operator != nil {
		operator = *req.Operator
	}

	var value float64
	if req.Value != nil {
		value = *req.Value
	}

	var floatKey string
	if req.FloatKey != nil {
		floatKey = *req.FloatKey
	}

	var choiceID uuid.UUID
	if req.ChoiceID != nil {
		choiceID, err = uuid.Parse(*req.ChoiceID)
		if err != nil {
			return nil, err
		}
	}

	row, err := s.q.CreateCondition(ctx, db.CreateConditionParams{
		Column1: uuid.New(),
		Column2: stepID,
		Column3: db.ConditionType(req.Type),
		Column4: metric,
		Column5: operator,
		Column6: value,
		Column7: floatKey,
		Column8: choiceID,
	})
	if err != nil {
		return nil, err
	}

	resp := &dto.ConditionCreatedResponse{
		ID:     row.ID.String(),
		StepID: row.StepID.String(),
		Type:   string(row.Type),
	}

	if row.Metric.Valid {
		m := string(row.Metric.ImpactMetric)
		resp.Metric = &m
	}
	if row.Operator.Valid {
		resp.Operator = ptrString(row.Operator)
	}
	if row.Value.Valid {
		v := row.Value.Float64
		resp.Value = &v
	}
	if row.FloatKey.Valid {
		resp.FloatKey = ptrString(row.FloatKey)
	}
	if row.ChoiceID != uuid.Nil {
		s := row.ChoiceID.String()
		resp.ChoiceID = &s
	}

	return resp, nil
}

func (s *ConditionService) DeleteCondition(ctx context.Context, id uuid.UUID) error {
	return s.q.DeleteCondition(ctx, id)
}

func ptrString(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}
	return &t.String
}
