package service

import (
	"project-app-inventory-restapi-golang-rahmadhany/dto"
	"project-app-inventory-restapi-golang-rahmadhany/model"
	"project-app-inventory-restapi-golang-rahmadhany/repository"
)

type WarehouseService interface {
	GetAll() ([]model.Warehouse, error)
	GetByID(id int) (*model.Warehouse, error)
	Create(input dto.CreateWarehouseRequest) (int, error)
	Update(id int, input dto.UpdateWarehouseRequest) error
	Delete(id int) error
}

type warehouseService struct {
	Repo repository.Repository
}

func NewWarehouseService(repo repository.Repository) WarehouseService {
	return &warehouseService{Repo: repo}
}

func (s *warehouseService) GetAll() ([]model.Warehouse, error) {
	return s.Repo.WarehouseRepo.GetAll()
}

func (s *warehouseService) GetByID(id int) (*model.Warehouse, error) {
	return s.Repo.WarehouseRepo.GetByID(id)
}

func (s *warehouseService) Create(input dto.CreateWarehouseRequest) (int, error) {
	return s.Repo.WarehouseRepo.Create(model.Warehouse{
		Name:        input.Name,
		Description: input.Description,
	})
}

func (s *warehouseService) Update(id int, input dto.UpdateWarehouseRequest) error {
	return s.Repo.WarehouseRepo.Update(model.Warehouse{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	})
}

func (s *warehouseService) Delete(id int) error {
	return s.Repo.WarehouseRepo.Delete(id)
}
