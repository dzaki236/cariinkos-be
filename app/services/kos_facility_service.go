package services

import (
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"
)

type KosFacilityService struct {
	repo    repositoryContracts.KosFacilityRepositoryContract
	kosRepo repositoryContracts.KosRepositoryContract
}

func NewKosFacilityService(repo repositoryContracts.KosFacilityRepositoryContract, kosRepo repositoryContracts.KosRepositoryContract) *KosFacilityService {
	return &KosFacilityService{repo: repo, kosRepo: kosRepo}
}

func (s *KosFacilityService) GetAllByKosID(kosID string) ([]models.Facility, error) {
	return s.repo.GetAllByKosID(kosID)
}

func (s *KosFacilityService) Attach(userID, kosID, facilityID string) error {
	if err := CheckKosOwner(userID, kosID, s.kosRepo); err != nil {
		return err
	}
	return s.repo.Attach(kosID, facilityID)
}

func (s *KosFacilityService) Detach(userID, kosID, facilityID string) error {
	if err := CheckKosOwner(userID, kosID, s.kosRepo); err != nil {
		return err
	}
	return s.repo.Detach(kosID, facilityID)
}

var _ serviceContracts.KosFacilityService = (*KosFacilityService)(nil)
