package dto

type StepResponse struct {
	ID       string           `json:"id"`
	Question string           `json:"question"`
	Context  *string          `json:"context,omitempty"`
	Hint     *string          `json:"hint,omitempty"`
	Choices  []ChoiceResponse `json:"choices"`
}
