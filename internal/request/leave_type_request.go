package request

type LevaeTypeRequest struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status"`
}
