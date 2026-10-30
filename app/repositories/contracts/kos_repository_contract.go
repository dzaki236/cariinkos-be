package contracts

import (
	"cariinkos/app/models"
)

type KosRepositoryContract interface {
	GetAllData(page, perPage int) ([]models.Kos, int64, error)
	GetByID(id string) (models.Kos, error)
	GetPublishedData(page, perPage int) ([]models.Kos, int64, error)
	GetPublishedByID(id string) (models.Kos, error)
	Create(kos *models.Kos) error
	Update(kos *models.Kos) error
	Delete(id string) error
}
