package service

import "project-app-inventory-restapi-golang-rahmadhany/repository"

type Service struct {
	UserService      UserService
	CategoryService  CategoryService
	RackService      RackService
	WarehouseService WarehouseService
	ProductService   ProductService
	SaleService      SaleService
}

func NewService(repo repository.Repository) Service {
	return Service{
		UserService:      NewUserService(repo),
		CategoryService:  NewCategoryService(repo),
		RackService:      NewRackService(repo),
		WarehouseService: NewWarehouseService(repo),
		ProductService:   NewProductService(repo),
		SaleService:      NewSaleService(repo),
	}
}
