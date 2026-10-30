package contracts

import "cariinkos/app/models"

type KosPhotoRepositoryContract interface {
	GetAllByKosID(kosID string) ([]models.KosPhoto, error)
	GetByID(kosID, photoID string) (models.KosPhoto, error)
	Create(photo *models.KosPhoto) error
	Update(photo *models.KosPhoto) error
	Delete(kosID, photoID string) error
}
