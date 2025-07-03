package dto

type CreateRackRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateRackRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
