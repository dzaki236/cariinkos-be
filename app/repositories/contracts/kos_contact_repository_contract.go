package contracts

import "cariinkos/app/models"

type KosContactRepositoryContract interface {
	GetByKosID(kosID string) (models.KosContact, error)
	Create(contact *models.KosContact) error
	Update(contact *models.KosContact) error
	Delete(kosID string) error
}
