package services

import (
	"cariinkos/app/dto"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"
)

type KosContactService struct {
	repo    repositoryContracts.KosContactRepositoryContract
	kosRepo repositoryContracts.KosRepositoryContract
}

func NewKosContactService(repo repositoryContracts.KosContactRepositoryContract, kosRepo repositoryContracts.KosRepositoryContract) *KosContactService {
	return &KosContactService{repo: repo, kosRepo: kosRepo}
}

func (s *KosContactService) GetByKosID(kosID string) (models.KosContact, error) {
	return s.repo.GetByKosID(kosID)
}

func (s *KosContactService) Create(userID, kosID string, request dto.CreateKosContactRequest) (models.KosContact, error) {
	if err := s.ensureOwner(userID, kosID); err != nil {
		return models.KosContact{}, err
	}
	contact := models.KosContact{KosID: kosID, ContactName: request.ContactName, Phone: request.Phone, Whatsapp: request.Whatsapp}
	if err := s.repo.Create(&contact); err != nil {
		return models.KosContact{}, err
	}
	return contact, nil
}

func (s *KosContactService) Update(userID, kosID string, request dto.UpdateKosContactRequest) (models.KosContact, error) {
	if err := s.ensureOwner(userID, kosID); err != nil {
		return models.KosContact{}, err
	}
	contact, err := s.repo.GetByKosID(kosID)
	if err != nil {
		return models.KosContact{}, err
	}
	contact.ContactName = request.ContactName
	contact.Phone = request.Phone
	contact.Whatsapp = request.Whatsapp
	if err := s.repo.Update(&contact); err != nil {
		return models.KosContact{}, err
	}
	return contact, nil
}

func (s *KosContactService) Delete(userID, kosID string) error {
	if err := s.ensureOwner(userID, kosID); err != nil {
		return err
	}
	return s.repo.Delete(kosID)
}

func (s *KosContactService) ensureOwner(userID, kosID string) error {
	kos, err := s.kosRepo.GetByID(kosID)
	if err != nil {
		return err
	}
	if kos.UserID != userID {
		return helpers.ErrDataNotFound
	}
	return nil
}

var _ serviceContracts.KosContactService = (*KosContactService)(nil)
