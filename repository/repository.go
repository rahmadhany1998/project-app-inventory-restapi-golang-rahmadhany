package repository

import (
	"database/sql"
)

type Repository struct {
	UserRepo      UserRepository
	CategoryRepo  CategoryRepository
	RackRepo      RackRepository
	WarehouseRepo WarehouseRepository
	ProductRepo   ProductRepository
	SaleRepo      SaleRepository
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		UserRepo:      NewUserRepository(db),
		CategoryRepo:  NewCategoryRepository(db),
		RackRepo:      NewRackRepository(db),
		WarehouseRepo: NewWarehouseRepository(db),
		ProductRepo:   NewProductRepository(db),
		SaleRepo:      NewSaleRepository(db),
	}
}
