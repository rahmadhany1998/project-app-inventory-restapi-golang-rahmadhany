package model

type Product struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	CategoryID     int    `json:"category_id"`
	RackID         int    `json:"rack_id"`
	WarehouseID    int    `json:"warehouse_id"`
	InventoryCount int    `json:"inventory_count"`
	RetailPrice    int    `json:"retail_price"`
	SellingPrice   int    `json:"selling_price"`
	Image          string `json:"image"`
}
