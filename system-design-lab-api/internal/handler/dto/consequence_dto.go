package dto

type CreateConsequenceRequest struct {
	ChoiceID string `json:"choiceId" validate:"required,uuid"`
	Type     string `json:"type" validate:"required"`
	Keys     string `json:"keys" validate:"required"`
	Value    bool   `json:"value"`
}

type UpdateConsequenceRequest struct {
	Type  string `json:"type" validate:"required"`
	Keys  string `json:"keys" validate:"required"`
	Value bool   `json:"value"`
}

type ConsequenceCreatedResponse struct {
	ID       string `json:"id"`
	ChoiceID string `json:"choiceId"`
	Type     string `json:"type"`
	Keys     string `json:"keys"`
	Value    bool   `json:"value"`
}
