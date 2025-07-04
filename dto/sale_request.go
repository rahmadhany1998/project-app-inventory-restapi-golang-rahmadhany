package dto

type CreateSaleRequest struct {
	ProductID int    `json:"product_id"`
	ItemSold  int    `json:"item_sold"`
	DateSale  string `json:"date_sale"`
}
