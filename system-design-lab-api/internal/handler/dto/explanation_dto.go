package dto

type CreateExplanationRequest struct {
	ChoiceID string `json:"choiceId" validate:"required,uuid"`
	Level    string `json:"level" validate:"required,oneof=beginner intermediate advanced"`
	Content  string `json:"content" validate:"required,min=1"`
}

type UpdateExplanationRequest struct {
	Level   string `json:"level" validate:"required,oneof=beginner intermediate advanced"`
	Content string `json:"content" validate:"required,min=1"`
}

type ExplanationCreatedResponse struct {
	ID       string `json:"id"`
	ChoiceID string `json:"choiceId"`
	Level    string `json:"level"`
	Content  string `json:"content"`
}
