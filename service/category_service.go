package service

import (
	"project-app-inventory-restapi-golang-rahmadhany/dto"
	"project-app-inventory-restapi-golang-rahmadhany/model"
	"project-app-inventory-restapi-golang-rahmadhany/repository"
)

type CategoryService interface {
	GetAll() ([]model.Category, error)
	GetByID(id int) (*model.Category, error)
	Create(input dto.CreateCategoryRequest) (int, error)
	Update(id int, input dto.UpdateCategoryRequest) error
	Delete(id int) error
}

type categoryService struct {
	Repo repository.Repository
}

func NewCategoryService(repo repository.Repository) CategoryService {
	return &categoryService{Repo: repo}
}

func (s *categoryService) GetAll() ([]model.Category, error) {
	return s.Repo.CategoryRepo.GetAll()
}

func (s *categoryService) GetByID(id int) (*model.Category, error) {
	return s.Repo.CategoryRepo.GetByID(id)
}

func (s *categoryService) Create(input dto.CreateCategoryRequest) (int, error) {
	return s.Repo.CategoryRepo.Create(model.Category{
		Name:        input.Name,
		Description: input.Description,
	})
}

func (s *categoryService) Update(id int, input dto.UpdateCategoryRequest) error {
	return s.Repo.CategoryRepo.Update(model.Category{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	})
}

func (s *categoryService) Delete(id int) error {
	return s.Repo.CategoryRepo.Delete(id)
}
