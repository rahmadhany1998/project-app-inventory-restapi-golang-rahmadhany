package dto

type CreateProductRequest struct {
	Name           string `json:"name" validate:"required"`
	CategoryID     int    `json:"category_id" validate:"numeric,required"`
	RackID         int    `json:"rack_id" validate:"numeric.required"`
	WarehouseID    int    `json:"warehouse_id" validate:"numeric,required"`
	InventoryCount int    `json:"inventory_count" validate:"numeric,required"`
	RetailPrice    int    `json:"retail_price" validate:"numeric,required"`
	SellingPrice   int    `json:"selling_price" validate:"numeric,required"`
	Image          string `json:"image" validate:"required"`
}

type UpdateProductRequest struct {
	Name           string `json:"name" validate:"required"`
	CategoryID     int    `json:"category_id" validate:"numeric,required"`
	RackID         int    `json:"rack_id" validate:"numeric,required"`
	WarehouseID    int    `json:"warehouse_id" validate:"numeric,min=6,required"`
	InventoryCount int    `json:"inventory_count" validate:"numeric,required"`
	RetailPrice    int    `json:"retail_price" validate:"numeric,required"`
	SellingPrice   int    `json:"selling_price" validate:"numeric,required"`
	Image          string `json:"image" validate:"required"`
}
