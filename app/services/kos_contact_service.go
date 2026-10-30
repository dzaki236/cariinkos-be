package services

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/google/uuid"
)

type KosContactService struct {
	repo repositoryContracts.KosContactRepositoryContract
}

func NewKosContactService(repo repositoryContracts.KosContactRepositoryContract) *KosContactService {
	return &KosContactService{repo: repo}
}

func (s *KosContactService) GetAllByKosID(kosID string) ([]models.KosContact, error) {
	return s.repo.GetAllByKosID(kosID)
}

func (s *KosContactService) Create(kosID string, request dto.CreateKosContactRequest) (models.KosContact, error) {
	contact := models.KosContact{ID: uuid.NewString(), KosID: kosID, Phone: request.Phone, Whatsapp: request.Whatsapp}
	if err := s.repo.Create(&contact); err != nil {
		return models.KosContact{}, err
	}
	return contact, nil
}

func (s *KosContactService) Update(kosID, contactID string, request dto.UpdateKosContactRequest) (models.KosContact, error) {
	contact, err := s.repo.GetByID(kosID, contactID)
	if err != nil {
		return models.KosContact{}, err
	}
	contact.Phone = request.Phone
	contact.Whatsapp = request.Whatsapp
	if err := s.repo.Update(&contact); err != nil {
		return models.KosContact{}, err
	}
	return contact, nil
}

func (s *KosContactService) Delete(kosID, contactID string) error {
	return s.repo.Delete(kosID, contactID)
}

var _ serviceContracts.KosContactService = (*KosContactService)(nil)
