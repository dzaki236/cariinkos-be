package contracts

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
)

type KosPhotoService interface {
	GetAllByKosID(kosID string) ([]models.KosPhoto, error)
	Create(kosID string, request dto.CreateKosPhotoRequest) (models.KosPhoto, error)
	Update(kosID, photoID string, request dto.UpdateKosPhotoRequest) (models.KosPhoto, error)
	Delete(kosID, photoID string) error
}
