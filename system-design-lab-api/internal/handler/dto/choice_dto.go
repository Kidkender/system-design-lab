package dto

type CreateChoiceRequest struct {
	StepID     string  `json:"stepId" validate:"required,uuid"`
	Label      string  `json:"label" validate:"required,min=1"`
	NextStepID *string `json:"nextStepId"`
	IsCorrect  bool    `json:"isCorrect"`
}

type ChoiceCreatedResponse struct {
	ID         string  `json:"id"`
	StepID     string  `json:"stepId"`
	Label      string  `json:"label"`
	NextStepID *string `json:"nextStepId"`
	IsCorrect  bool    `json:"isCorrect"`
}
