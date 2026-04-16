package dto

type CreateConditionRequest struct {
	StepID   string   `json:"stepId" validate:"required,uuid"`
	Type     string   `json:"type" validate:"required,oneof=metric flag choice"`
	Metric   *string  `json:"metric"`
	Operator *string  `json:"operator"`
	Value    *float64 `json:"value"`
	FloatKey *string  `json:"floatKey"`
	ChoiceID *string  `json:"choiceId"`
}

type ConditionCreatedResponse struct {
	ID       string   `json:"id"`
	StepID   string   `json:"stepId"`
	Type     string   `json:"type"`
	Metric   *string  `json:"metric,omitempty"`
	Operator *string  `json:"operator,omitempty"`
	Value    *float64 `json:"value,omitempty"`
	FloatKey *string  `json:"floatKey,omitempty"`
	ChoiceID *string  `json:"choiceId,omitempty"`
}
