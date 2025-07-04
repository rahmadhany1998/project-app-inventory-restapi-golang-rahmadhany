package handler

import "project-app-inventory-restapi-golang-rahmadhany/service"

type Handler struct {
	UserHandler      UserHandler
	CategoryHandler  CategoryHandler
	RackHandler      RackHandler
	WarehouseHandler WarehouseHandler
}

func NewHandler(s service.Service) Handler {
	return Handler{
		UserHandler:      NewUserHandler(s),
		CategoryHandler:  NewCategoryHandler(s),
		RackHandler:      NewRackHandler(s),
		WarehouseHandler: NewWarehouseHandler(s),
	}
}
