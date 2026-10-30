package contracts

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
)

type KosContactService interface {
	GetAllByKosID(kosID string) ([]models.KosContact, error)
	Create(kosID string, request dto.CreateKosContactRequest) (models.KosContact, error)
	Update(kosID, contactID string, request dto.UpdateKosContactRequest) (models.KosContact, error)
	Delete(kosID, contactID string) error
}
