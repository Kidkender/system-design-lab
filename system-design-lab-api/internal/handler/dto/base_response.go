package dto

type PaginationResponse[T any] struct {
	Data       []T   `json:"data"`
	Total      int32 `json:"total"`
	Page       int32 `json:"page"`
	Limit      int32 `json:"limit"`
	TotalPages int32 `json:"totalPages"`
}

// Concrete types for swaggo (does not support generics)

type ScenarioPaginationResponse struct {
	Data       []ScenarioResponse `json:"data"`
	Total      int32              `json:"total"`
	Page       int32              `json:"page"`
	Limit      int32              `json:"limit"`
	TotalPages int32              `json:"totalPages"`
}

type StepPaginationResponse struct {
	Data       []StepResponse `json:"data"`
	Total      int32          `json:"total"`
	Page       int32          `json:"page"`
	Limit      int32          `json:"limit"`
	TotalPages int32          `json:"totalPages"`
}
