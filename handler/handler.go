package handler

import (
	"project-app-inventory-restapi-golang-rahmadhany/service"
	"project-app-inventory-restapi-golang-rahmadhany/utils"
)

type Handler struct {
	UserHandler      UserHandler
	CategoryHandler  CategoryHandler
	RackHandler      RackHandler
	WarehouseHandler WarehouseHandler
	ProductHandler   ProductHandler
	SaleHandler      SaleHandler
}

func NewHandler(s service.Service, cfg utils.Configuration) Handler {
	return Handler{
		UserHandler:      NewUserHandler(s, cfg),
		CategoryHandler:  NewCategoryHandler(s, cfg),
		RackHandler:      NewRackHandler(s, cfg),
		WarehouseHandler: NewWarehouseHandler(s, cfg),
		ProductHandler:   NewProductHandler(s, cfg),
		SaleHandler:      NewSaleHandler(s, cfg),
	}
}
