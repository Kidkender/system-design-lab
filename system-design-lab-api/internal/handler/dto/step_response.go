package dto

type StepResponse struct {
	ID       string           `json:"id"`
	Question string           `json:"question"`
	Context  *string          `json:"context"`
	Choices  []ChoiceResponse `json:"choices"`
}
