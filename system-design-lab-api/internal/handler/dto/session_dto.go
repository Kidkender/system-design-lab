package dto

type StartSessionRequest struct {
	UserID     string `json:"userId" validate:"required,uuid"`
	ScenarioID string `json:"scenarioId" validate:"required,uuid"`
}

type SessionResponse struct {
	ID          string       `json:"id"`
	ScenarioID  string       `json:"scenarioId"`
	CurrentStep StepResponse `json:"currentStep"`
	Metrics     map[string]float64 `json:"metrics"`
	Flags       map[string]bool    `json:"flags"`
	Status      string             `json:"status"`
}

type SubmitChoiceRequest struct {
	ChoiceID string `json:"choiceId" validate:"required,uuid"`
}

type SubmitChoiceResponse struct {
	IsCorrect   bool               `json:"isCorrect"`
	Feedback    string             `json:"feedback"`
	Metrics     map[string]float64 `json:"metrics"`
	NextStep    *StepResponse      `json:"nextStep"`
	IsCompleted bool               `json:"isCompleted"`
}
