package dto

type CreateScenarioRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	StartStepID *string `json:"start_step_id"`
	Difficulty  string  `json:"difficulty"`
}

type UpdateStartStepRequest struct {
	StartStepID *string `json:"start_step_id"`
}
