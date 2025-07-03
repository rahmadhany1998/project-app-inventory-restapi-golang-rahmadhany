package service

import "project-app-inventory-restapi-golang-rahmadhany/repository"

type Service struct {
	UserService UserService
}

func NewService(repo repository.Repository) Service {
	return Service{
		UserService: NewUserService(repo),
	}
}
