package dto

type CreateImpactRequest struct {
	ChoiceID string  `json:"choiceId" validate:"required,uuid"`
	Metric   string  `json:"metric" validate:"required,oneof=latency cost scalability"`
	Type     string  `json:"type" validate:"required,oneof=add multiply set"`
	Value    float64 `json:"value"`
}

type UpdateImpactRequest struct {
	Metric string  `json:"metric" validate:"required,oneof=latency cost scalability"`
	Type   string  `json:"type" validate:"required,oneof=add multiply set"`
	Value  float64 `json:"value"`
}

type ImpactCreatedResponse struct {
	ID       string  `json:"id"`
	ChoiceID string  `json:"choiceId"`
	Metric   string  `json:"metric"`
	Type     string  `json:"type"`
	Value    float64 `json:"value"`
}
