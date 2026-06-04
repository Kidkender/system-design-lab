package dto

type CreateScenarioRequest struct {
	Title       string  `json:"title" validate:"required,min=3"`
	Description string  `json:"description" validate:"max=500"`
	StartStepID *string `json:"start_step_id"`
	Difficulty  string  `json:"difficulty" validate:"oneof=easy medium hard"`
}

type UpdateStartStepRequest struct {
	StartStepID *string `json:"start_step_id" validate:"required"`
}

type ScenarioFilter struct {
	Difficulty *string `json:"difficulty" validate:"oneof=easy medium hard"`
	Page       int
	Limit      int
}
