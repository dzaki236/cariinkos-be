package contracts

import (
	"cariinkos/app/models"
)

type KosRepositoryContract interface {
	GetAllData() ([]models.Kos, error)
	GetByID(id string) (models.Kos, error)
	Create(kos *models.Kos) error
	Update(kos *models.Kos) error
	Delete(id string) error
}
