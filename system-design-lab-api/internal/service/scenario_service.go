package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kidkender/system-design-lab/internal/db"

	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type ScenarioService struct {
	q *db.Queries
}

func NewScenarioService(q *db.Queries) *ScenarioService {
	return &ScenarioService{q: q}
}

func (s *ScenarioService) CreateScenario(ctx context.Context, req *dto.CreateScenarioRequest) {
	scenario, err := s.q.CreateScenario(
		ctx,
		db.CreateScenarioParams{
			ID:          uuid.New(),
			Title:       req.Title,
			Description: pgtype.Text{String: req.Description, Valid: true},
			Difficulty:  db.DifficultyLevel(req.Difficulty),
		},
	)
	if err != nil {
		slog.Error("Failed to create scenario", "error", err)
	}

	slog.Info("Scenario %s created", "id", scenario.ID)
}

func (s *ScenarioService) GetScenario(ctx context.Context, id uuid.UUID) (*dto.ScenarioResponse, error) {
	sc, err := s.q.GetScenario(ctx, id)
	if err != nil {
		return nil, err
	}

	steps, err := s.q.GetStepsByScenario(ctx, id)
	if err != nil {
		return nil, err
	}

	stepIDs := make([]uuid.UUID, 0, len(steps))
	for _, st := range steps {
		stepIDs = append(stepIDs, st.ID)
	}

	choices, err := s.q.GetChoicesByStepIDs(ctx, stepIDs)
	if err != nil {
		return nil, err
	}

	choiceIDs := make([]uuid.UUID, 0, len(choices))
	for _, c := range choices {
		choiceIDs = append(choiceIDs, c.ID)
	}

	impacts, _ := s.q.GetImpactsByChoiceIDs(ctx, choiceIDs)
	explanations, _ := s.q.GetExplanationsByChoiceIDs(ctx, choiceIDs)
	consequences, _ := s.q.GetConsequencesByChoiceIDs(ctx, choiceIDs)

	choiceMap := map[uuid.UUID][]db.GetChoicesByStepIDsRow{}
	for _, c := range choices {
		choiceMap[c.StepID] = append(choiceMap[c.StepID], c)
	}

	impactMap := map[uuid.UUID][]db.GetImpactsByChoiceIDsRow{}
	for _, i := range impacts {
		impactMap[i.ChoiceID] = append(impactMap[i.ChoiceID], i)
	}

	explainMap := map[uuid.UUID]map[string]string{}
	for _, e := range explanations {
		if explainMap[e.ChoiceID] == nil {
			explainMap[e.ChoiceID] = map[string]string{}
		}

		explainMap[e.ChoiceID][string(e.Level)] = e.Content

	}

	consequenceMap := map[uuid.UUID][]db.GetConsequencesByChoiceIDsRow{}
	for _, c := range consequences {
		consequenceMap[c.ChoiceID] = append(consequenceMap[c.ChoiceID], c)
	}

	resp := &dto.ScenarioResponse{
		ID:          sc.ID.String(),
		Title:       sc.Title,
		Description: sc.Description.String,
		Steps:       []dto.StepResponse{},
	}

	for _, st := range steps {
		stepResp := dto.StepResponse{
			ID:       st.ID.String(),
			Question: st.Question,
			Context:  &st.Context.String,
			Choices:  []dto.ChoiceResponse{},
		}

		for _, c := range choiceMap[st.ID] {
			nextStepID := c.NextStepID.String()
			ch := dto.ChoiceResponse{
				ID:         c.ID.String(),
				Label:      c.Label,
				NextStepID: &nextStepID,
				IsCorrect:  c.IsCorrect,
			}

			// impacts
			for _, i := range impactMap[c.ID] {
				ch.Impacts = append(ch.Impacts, dto.ImpactResponse{
					Metric: string(i.Metric),
					Type:   string(i.Type),
					Value:  i.Value,
				})
			}

			// consequences
			for _, cs := range consequenceMap[c.ID] {
				ch.Consequences = append(ch.Consequences, dto.ConsequencesResponse{
					Key:   cs.Keys,
					Value: cs.Value,
				})
			}

			// explanations
			ch.Explanations = explainMap[c.ID]
			stepResp.Choices = append(stepResp.Choices, ch)

		}
		resp.Steps = append(resp.Steps, stepResp)
	}

	return resp, nil

}

func (s *ScenarioService) GetScenarios(ctx context.Context) ([]dto.ScenarioResponse, error) {
	scenarios, err := s.q.GetScenarios(ctx)
	if err != nil {
		return nil, err
	}

	resp := []dto.ScenarioResponse{}
	for _, sc := range scenarios {
		resp = append(resp, dto.ScenarioResponse{
			ID:          sc.ID.String(),
			Title:       sc.Title,
			Description: sc.Description.String,
		})
	}
	return resp, nil
}

func (s *ScenarioService) GetScenariosPaginated(
	ctx context.Context, page int32, limit int32,
) (*dto.PaginationResponse[dto.ScenarioResponse], error) {
	if page < 0 {
		page = 1
	}

	if limit < 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	scenarios, err := s.q.ListScenariosPaginated(ctx, db.ListScenariosPaginatedParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})

	if err != nil {
		return nil, err
	}

	data := make([]dto.ScenarioResponse, 0, len(scenarios))

	for _, sc := range scenarios {
		data = append(data, dto.ScenarioResponse{
			ID:          sc.ID.String(),
			Title:       sc.Title,
			Description: sc.Description.String,
			// Steps: s,
		})
	}
	totalPages := (int32(len(scenarios)) + limit - 1) / limit

	return &dto.PaginationResponse[dto.ScenarioResponse]{
		Data:       data,
		Total:      int32(len(scenarios)),
		Page:       page,
		Limit:      limit,
		TotalPages: int32(totalPages),
	}, nil

}
