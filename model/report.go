package model

type Report struct {
	TotalProduct int `json:"total_product"`
	TotalSale    int `json:"total_sale"`
	TotalRevenue int `json:"total_revenue"`
}
