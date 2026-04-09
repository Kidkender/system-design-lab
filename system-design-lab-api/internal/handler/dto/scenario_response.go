package dto

type ImpactResponse struct {
	Metric string  `json:"metric"`
	Type   string  `json:"type"`
	Value  float64 `json:"value"`
}

type ConsequencesResponse struct {
	Key   string `json:"key"`
	Value bool   `json:"value"`
}

type ChoiceResponse struct {
	ID           string                 `json:"id"`
	Label        string                 `json:"label"`
	NextStepID   *string                `json:"nextStepId"`
	IsCorrect    bool                   `json:"isCorrect"`
	Impacts      []ImpactResponse       `json:"impacts"`
	Consequences []ConsequencesResponse `json:"consequences"`
	Explanations map[string]string      `json:"explanations"`
}

type StepResponse struct {
	ID       string           `json:"id"`
	Question string           `json:"question"`
	Context  *string          `json:"context"`
	Choices  []ChoiceResponse `json:"choices"`
}

type ScenarioResponse struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Steps       []StepResponse `json:"steps"`
}
