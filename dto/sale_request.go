package dto

type CreateSaleRequest struct {
	ProductID int    `json:"product_id" validate:"numeric,required"`
	ItemSold  int    `json:"item_sold" validate:"numeric,min=1,required"`
	DateSale  string `json:"date_sale" validate:"required"`
}
