package service

import (
	"math"
	"project-app-inventory-restapi-golang-rahmadhany/dto"
	"project-app-inventory-restapi-golang-rahmadhany/model"
	"project-app-inventory-restapi-golang-rahmadhany/repository"
)

type SaleService interface {
	GetAll(page, limit int) ([]model.Sale, int, int, error)
	GetByID(id int) (*model.Sale, error)
	Create(input dto.CreateSaleRequest) (int, error)
	GetReportSummaryByDate(start, end string) (*model.Report, error)
}

type saleService struct {
	Repo repository.Repository
}

func NewSaleService(repo repository.Repository) SaleService {
	return &saleService{Repo: repo}
}

func (s *saleService) GetAll(page, limit int) ([]model.Sale, int, int, error) {
	if page < 1 {
		page = 1
	}

	totalRecords, err := s.Repo.SaleRepo.CountAll()
	if err != nil {
		return nil, 0, 0, err
	}
	sales, err := s.Repo.SaleRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	return sales, totalRecords, totalPages, nil
}

func (s *saleService) GetByID(id int) (*model.Sale, error) {
	return s.Repo.SaleRepo.GetByID(id)
}

func (s *saleService) Create(input dto.CreateSaleRequest) (int, error) {
	product, err := s.Repo.ProductRepo.GetByID(input.ProductID)
	if err != nil {
		return 0, err
	}

	//calculate total bill
	totalbill := input.ItemSold * product.SellingPrice

	return s.Repo.SaleRepo.Create(model.Sale{
		ProductID: input.ProductID,
		ItemSold:  input.ItemSold,
		TotalBill: totalbill,
		DateSale:  input.DateSale,
	})
}

func (s *saleService) GetReportSummaryByDate(start, end string) (*model.Report, error) {
	return s.Repo.SaleRepo.GetReportSummaryByDate(start, end)
}
