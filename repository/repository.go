package repository

import (
	"database/sql"
)

type Repository struct {
	UserRepo     UserRepository
	CategoryRepo CategoryRepository
	RackRepo     RackRepository
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		UserRepo:     NewUserRepository(db),
		CategoryRepo: NewCategoryRepository(db),
		RackRepo:     NewRackRepository(db),
	}
}
