package service

import (
	"math"
	"project-app-inventory-restapi-golang-rahmadhany/dto"
	"project-app-inventory-restapi-golang-rahmadhany/model"
	"project-app-inventory-restapi-golang-rahmadhany/repository"
)

type UserService interface {
	GetAll(page, limit int) ([]model.User, int, int, error)
	GetByID(id int) (*model.User, error)
	Create(input dto.CreateUserRequest) (int, error)
	Update(id int, input dto.UpdateUserRequest) error
	Delete(id int) error
}

type userService struct {
	Repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{Repo: repo}
}

func (s *userService) GetAll(page, limit int) ([]model.User, int, int, error) {
	if page < 1 {
		page = 1
	}

	totalRecords, err := s.Repo.UserRepo.CountAll()
	if err != nil {
		return nil, 0, 0, err
	}
	users, err := s.Repo.UserRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	return users, totalRecords, totalPages, nil
}

func (s *userService) GetByID(id int) (*model.User, error) {
	return s.Repo.UserRepo.GetByID(id)
}

func (s *userService) Create(input dto.CreateUserRequest) (int, error) {
	return s.Repo.UserRepo.Create(model.User{
		Name:     input.Name,
		Email:    input.Email,
		Role:     input.Role,
		Password: input.Password,
		Status:   input.Status,
	})
}

func (s *userService) Update(id int, input dto.UpdateUserRequest) error {
	return s.Repo.UserRepo.Update(model.User{
		ID:       id,
		Name:     input.Name,
		Email:    input.Email,
		Role:     input.Role,
		Password: input.Password,
		Status:   input.Status,
	})
}

func (s *userService) Delete(id int) error {
	return s.Repo.UserRepo.Delete(id)
}
