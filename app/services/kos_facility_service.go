package services

import (
	"cariinkos/app/helpers"
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
	if err := s.ensureOwner(userID, kosID); err != nil {
		return err
	}
	return s.repo.Attach(kosID, facilityID)
}

func (s *KosFacilityService) Detach(userID, kosID, facilityID string) error {
	if err := s.ensureOwner(userID, kosID); err != nil {
		return err
	}
	return s.repo.Detach(kosID, facilityID)
}

func (s *KosFacilityService) ensureOwner(userID, kosID string) error {
	kos, err := s.kosRepo.GetByID(kosID)
	if err != nil {
		return err
	}
	if kos.UserID != userID {
		return helpers.ErrDataNotFound
	}
	return nil
}

var _ serviceContracts.KosFacilityService = (*KosFacilityService)(nil)
