package service

import (
	"math"
	"project-app-inventory-restapi-golang-rahmadhany/dto"
	"project-app-inventory-restapi-golang-rahmadhany/model"
	"project-app-inventory-restapi-golang-rahmadhany/repository"
)

type CategoryService interface {
	GetAll(page, limit int) ([]model.Category, int, int, error)
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

func (s *categoryService) GetAll(page, limit int) ([]model.Category, int, int, error) {
	if page < 1 {
		page = 1
	}
	totalRecords, err := s.Repo.CategoryRepo.CountAll()
	if err != nil {
		return nil, 0, 0, err
	}
	categories, err := s.Repo.CategoryRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	return categories, totalRecords, totalPages, nil
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
