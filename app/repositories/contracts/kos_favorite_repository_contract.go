package contracts

import "cariinkos/app/models"

type KosFavoriteRepository interface {
	GetByUser(userID string) ([]models.KosFavorite, error)
	Find(userID, kosID string) (models.KosFavorite, error)
	Create(favorite *models.KosFavorite) error
	Update(favorite *models.KosFavorite) error
	Delete(userID, kosID string) error
}
