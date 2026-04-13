package dto

type CreateStepRequest struct {
	ScenarioID string  `json:"scenarioID"`
	Question   string  `json:"question" validate:"required,min=3"`
	Context    *string `json:"context" validate:"max=500"`
	OrderIndex int32   `json:"orderIndex"`
}
