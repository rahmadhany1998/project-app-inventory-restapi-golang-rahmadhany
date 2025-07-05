package service

import (
	"math"
	"project-app-inventory-restapi-golang-rahmadhany/dto"
	"project-app-inventory-restapi-golang-rahmadhany/model"
	"project-app-inventory-restapi-golang-rahmadhany/repository"
)

type RackService interface {
	GetAll(page, limit int) ([]model.Rack, int, int, error)
	GetByID(id int) (*model.Rack, error)
	Create(input dto.CreateRackRequest) (int, error)
	Update(id int, input dto.UpdateRackRequest) error
	Delete(id int) error
}

type rackService struct {
	Repo repository.Repository
}

func NewRackService(repo repository.Repository) RackService {
	return &rackService{Repo: repo}
}

func (s *rackService) GetAll(page, limit int) ([]model.Rack, int, int, error) {
	if page < 1 {
		page = 1
	}

	totalRecords, err := s.Repo.RackRepo.CountAll()
	if err != nil {
		return nil, 0, 0, err
	}
	racks, err := s.Repo.RackRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	return racks, totalRecords, totalPages, nil
}

func (s *rackService) GetByID(id int) (*model.Rack, error) {
	return s.Repo.RackRepo.GetByID(id)
}

func (s *rackService) Create(input dto.CreateRackRequest) (int, error) {
	return s.Repo.RackRepo.Create(model.Rack{
		Name:        input.Name,
		Description: input.Description,
	})
}

func (s *rackService) Update(id int, input dto.UpdateRackRequest) error {
	return s.Repo.RackRepo.Update(model.Rack{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	})
}

func (s *rackService) Delete(id int) error {
	return s.Repo.RackRepo.Delete(id)
}
