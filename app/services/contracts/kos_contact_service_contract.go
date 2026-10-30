package contracts

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
)

type KosContactService interface {
	GetByKosID(kosID string) (models.KosContact, error)
	Create(kosID string, request dto.CreateKosContactRequest) (models.KosContact, error)
	Update(kosID string, request dto.UpdateKosContactRequest) (models.KosContact, error)
	Delete(kosID string) error
}
