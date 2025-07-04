package dto

type CreateWarehouseRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateWarehouseRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
