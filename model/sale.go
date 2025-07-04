package model

type Sale struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	ItemSold  int    `json:"item_sold"`
	TotalBill int    `json:"total_bill"`
	DateSale  string `json:"date_sale"`
}
