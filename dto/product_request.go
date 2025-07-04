package dto

type CreateProductRequest struct {
	Name           string `json:"name"`
	CategoryID     int    `json:"category_id"`
	RackID         int    `json:"rack_id"`
	WarehouseID    int    `json:"warehouse_id"`
	InventoryCount int    `json:"inventory_count"`
	RetailPrice    int    `json:"retail_price"`
	SellingPrice   int    `json:"selling_price"`
	Image          string `json:"image"`
}

type UpdateProductRequest struct {
	Name           string `json:"name"`
	CategoryID     int    `json:"category_id"`
	RackID         int    `json:"rack_id"`
	WarehouseID    int    `json:"warehouse_id"`
	InventoryCount int    `json:"inventory_count"`
	RetailPrice    int    `json:"retail_price"`
	SellingPrice   int    `json:"selling_price"`
	Image          string `json:"image"`
}
