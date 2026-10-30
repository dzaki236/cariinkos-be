package services

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"
)

type KosContactService struct {
	repo repositoryContracts.KosContactRepositoryContract
}

func NewKosContactService(repo repositoryContracts.KosContactRepositoryContract) *KosContactService {
	return &KosContactService{repo: repo}
}

func (s *KosContactService) GetByKosID(kosID string) (models.KosContact, error) {
	return s.repo.GetByKosID(kosID)
}

func (s *KosContactService) Create(kosID string, request dto.CreateKosContactRequest) (models.KosContact, error) {
	contact := models.KosContact{KosID: kosID, ContactName: request.ContactName, Phone: request.Phone, Whatsapp: request.Whatsapp}
	if err := s.repo.Create(&contact); err != nil {
		return models.KosContact{}, err
	}
	return contact, nil
}

func (s *KosContactService) Update(kosID string, request dto.UpdateKosContactRequest) (models.KosContact, error) {
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

func (s *KosContactService) Delete(kosID string) error {
	return s.repo.Delete(kosID)
}

var _ serviceContracts.KosContactService = (*KosContactService)(nil)
