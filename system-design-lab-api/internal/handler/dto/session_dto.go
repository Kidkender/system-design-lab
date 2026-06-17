package dto

type StartSessionRequest struct {
	UserID     string `json:"userId" validate:"required,uuid"`
	ScenarioID string `json:"scenarioId" validate:"required,uuid"`
	Mode       string `json:"mode" validate:"omitempty,oneof=normal interview"`
}

type SessionResponse struct {
	ID               string             `json:"id"`
	ScenarioID       string             `json:"scenarioId"`
	CurrentStep      StepResponse       `json:"currentStep"`
	Metrics          map[string]float64 `json:"metrics"`
	Flags            map[string]bool    `json:"flags"`
	Status           string             `json:"status"`
	Mode             string             `json:"mode"`
	TimeLimitSeconds *int32             `json:"timeLimitSeconds,omitempty"`
	TimeElapsedSecs  int64              `json:"timeElapsedSeconds"`
}

type SubmitChoiceRequest struct {
	ChoiceID string `json:"choiceId" validate:"required,uuid"`
}

type UserChoiceSummary struct {
	StepID       string            `json:"stepId"`
	Question     string            `json:"question"`
	ChoiceID     string            `json:"choiceId"`
	Label        string            `json:"label"`
	IsCorrect    bool              `json:"isCorrect"`
	Explanations map[string]string `json:"explanations"`
}

type SessionSummaryResponse struct {
	ID         string              `json:"id"`
	ScenarioID string              `json:"scenarioId"`
	Status     string              `json:"status"`
	Metrics    map[string]float64  `json:"metrics"`
	Flags      map[string]bool     `json:"flags"`
	Choices    []UserChoiceSummary `json:"choices"`
	CreatedAt  string              `json:"createdAt"`
}

type SubmitChoiceResponse struct {
	IsCorrect    bool               `json:"isCorrect"`
	Feedback     string             `json:"feedback"`
	Explanations map[string]string  `json:"explanations"`
	Metrics      map[string]float64 `json:"metrics"`
	NextStep     *StepResponse      `json:"nextStep"`
	IsCompleted  bool               `json:"isCompleted"`
}

type UserSessionListItem struct {
	ID          string  `json:"id"`
	ScenarioID  string  `json:"scenarioId"`
	Status      string  `json:"status"`
	Mode        string  `json:"mode"`
	Score       float64 `json:"score"`
	CreatedAt   string  `json:"createdAt"`
	CompletedAt *string `json:"completedAt,omitempty"`
}

type LeaderboardEntry struct {
	Rank           int     `json:"rank"`
	SessionID      string  `json:"sessionId"`
	UserID         string  `json:"userId"`
	Username       string  `json:"username"`
	Score          float64 `json:"score"`
	TotalChoices   int32   `json:"totalChoices"`
	CorrectChoices int32   `json:"correctChoices"`
	CreatedAt      string  `json:"createdAt"`
	CompletedAt    string  `json:"completedAt"`
}

type UserProgressItem struct {
	ScenarioID      string  `json:"scenarioId"`
	Title           string  `json:"title"`
	Difficulty      string  `json:"difficulty"`
	Attempts        int32   `json:"attempts"`
	Completions     int32   `json:"completions"`
	BestScore       float64 `json:"bestScore"`
	LastCompletedAt *string `json:"lastCompletedAt,omitempty"`
}
