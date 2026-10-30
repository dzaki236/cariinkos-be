package contracts

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
)

type KosService interface {
	GetAllData(page, perPage int) ([]models.Kos, int64, error)
	GetByID(id string) (models.Kos, error)
	Create(userID string, request dto.CreateKosRequest) (models.Kos, error)
	Update(id string, request dto.UpdateKosRequest) (models.Kos, error)
	Delete(id string) error
}
