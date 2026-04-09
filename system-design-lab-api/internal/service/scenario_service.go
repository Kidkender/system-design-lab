package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/db"

	"github.com/kidkender/system-design-lab/internal/handler/dto"
)

type ScenarioService struct {
	q *db.Queries
}

func NewScenarioService(q *db.Queries) *ScenarioService {
	return &ScenarioService{q: q}
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
