package contracts

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
)

type KosContactService interface {
	GetByKosID(kosID string) (models.KosContact, error)
	Create(userID string, kosID string, request dto.CreateKosContactRequest) (models.KosContact, error)
	Update(userID string, kosID string, request dto.UpdateKosContactRequest) (models.KosContact, error)
	Delete(userID string, kosID string) error
}
