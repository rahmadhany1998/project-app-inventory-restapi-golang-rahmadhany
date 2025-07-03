package service

import "project-app-inventory-restapi-golang-rahmadhany/repository"

type Service struct {
	UserService     UserService
	CategoryService CategoryService
	RackService     RackService
}

func NewService(repo repository.Repository) Service {
	return Service{
		UserService:     NewUserService(repo),
		CategoryService: NewCategoryService(repo),
		RackService:     NewRackService(repo),
	}
}
