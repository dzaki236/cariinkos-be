package contracts

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
)

type KosFavoriteService interface {
	GetByUser(userID string) ([]models.KosFavorite, error)
	Create(userID, kosID string, request dto.UpdateFavoriteRequest) (models.KosFavorite, error)
	Update(userID, kosID string, request dto.UpdateFavoriteRequest) (models.KosFavorite, error)
	Delete(userID, kosID string) error
}
