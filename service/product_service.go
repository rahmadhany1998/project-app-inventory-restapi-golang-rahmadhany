package service

import (
	"math"
	"project-app-inventory-restapi-golang-rahmadhany/dto"
	"project-app-inventory-restapi-golang-rahmadhany/model"
	"project-app-inventory-restapi-golang-rahmadhany/repository"
)

type ProductService interface {
	GetAll(page, limit int) ([]model.Product, int, int, error)
	GetByID(id int) (*model.Product, error)
	Create(input dto.CreateProductRequest) (int, error)
	Update(id int, input dto.UpdateProductRequest) error
	Delete(id int) error
}

type productService struct {
	Repo repository.Repository
}

func NewProductService(repo repository.Repository) ProductService {
	return &productService{Repo: repo}
}

func (s *productService) GetAll(page, limit int) ([]model.Product, int, int, error) {
	if page < 1 {
		page = 1
	}

	totalRecords, err := s.Repo.ProductRepo.CountAll()
	if err != nil {
		return nil, 0, 0, err
	}
	products, err := s.Repo.ProductRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	return products, totalRecords, totalPages, nil
}

func (s *productService) GetByID(id int) (*model.Product, error) {
	return s.Repo.ProductRepo.GetByID(id)
}

func (s *productService) Create(input dto.CreateProductRequest) (int, error) {
	return s.Repo.ProductRepo.Create(model.Product{
		Name:           input.Name,
		CategoryID:     input.CategoryID,
		RackID:         input.RackID,
		WarehouseID:    input.WarehouseID,
		InventoryCount: input.InventoryCount,
		RetailPrice:    input.RetailPrice,
		SellingPrice:   input.SellingPrice,
		Image:          input.Image,
	})
}

func (s *productService) Update(id int, input dto.UpdateProductRequest) error {
	return s.Repo.ProductRepo.Update(model.Product{
		ID:             id,
		Name:           input.Name,
		CategoryID:     input.CategoryID,
		RackID:         input.RackID,
		WarehouseID:    input.WarehouseID,
		InventoryCount: input.InventoryCount,
		RetailPrice:    input.RetailPrice,
		SellingPrice:   input.SellingPrice,
		Image:          input.Image,
	})
}

func (s *productService) Delete(id int) error {
	return s.Repo.ProductRepo.Delete(id)
}
