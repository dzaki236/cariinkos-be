package services

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/google/uuid"
)

type FacilityService struct {
	repo repositoryContracts.FacilityRepositoryContract
}

func NewFacilityService(repo repositoryContracts.FacilityRepositoryContract) *FacilityService {
	return &FacilityService{repo: repo}
}

func (s *FacilityService) GetAllData() ([]models.Facility, error) {
	return s.repo.GetAllData()
}

func (s *FacilityService) GetByID(id string) (models.Facility, error) {
	return s.repo.GetByID(id)
}

func (s *FacilityService) Create(request dto.CreateFacilityRequest) (models.Facility, error) {
	facility := models.Facility{
		ID:   uuid.NewString(),
		Name: request.Name,
		Icon: request.Icon,
	}
	if err := s.repo.Create(&facility); err != nil {
		return models.Facility{}, err
	}
	return facility, nil
}

func (s *FacilityService) Update(id string, request dto.UpdateFacilityRequest) (models.Facility, error) {
	facility, err := s.repo.GetByID(id)
	if err != nil {
		return models.Facility{}, err
	}
	facility.Name = request.Name
	facility.Icon = request.Icon
	if err := s.repo.Update(&facility); err != nil {
		return models.Facility{}, err
	}
	return facility, nil
}

func (s *FacilityService) Delete(id string) error {
	return s.repo.Delete(id)
}

var _ serviceContracts.FacilityService = (*FacilityService)(nil)
