package services

import (
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"
)

type KosFacilityService struct {
	repo repositoryContracts.KosFacilityRepositoryContract
}

func NewKosFacilityService(repo repositoryContracts.KosFacilityRepositoryContract) *KosFacilityService {
	return &KosFacilityService{repo: repo}
}

func (s *KosFacilityService) GetAllByKosID(kosID string) ([]models.Facility, error) {
	return s.repo.GetAllByKosID(kosID)
}

func (s *KosFacilityService) Attach(kosID, facilityID string) error {
	return s.repo.Attach(kosID, facilityID)
}

func (s *KosFacilityService) Detach(kosID, facilityID string) error {
	return s.repo.Detach(kosID, facilityID)
}

var _ serviceContracts.KosFacilityService = (*KosFacilityService)(nil)
