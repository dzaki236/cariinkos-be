package contracts

import "cariinkos/app/models"

type KosContactRepositoryContract interface {
	GetAllByKosID(kosID string) ([]models.KosContact, error)
	GetByID(kosID, contactID string) (models.KosContact, error)
	Create(contact *models.KosContact) error
	Update(contact *models.KosContact) error
	Delete(kosID, contactID string) error
}
