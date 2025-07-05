package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type Repository struct {
	UserRepo      UserRepository
	CategoryRepo  CategoryRepository
	RackRepo      RackRepository
	WarehouseRepo WarehouseRepository
	ProductRepo   ProductRepository
	SaleRepo      SaleRepository
}

func NewRepository(db *sql.DB, log *zap.Logger) Repository {
	return Repository{
		UserRepo:      NewUserRepository(db, log),
		CategoryRepo:  NewCategoryRepository(db, log),
		RackRepo:      NewRackRepository(db, log),
		WarehouseRepo: NewWarehouseRepository(db, log),
		ProductRepo:   NewProductRepository(db, log),
		SaleRepo:      NewSaleRepository(db, log),
	}
}
