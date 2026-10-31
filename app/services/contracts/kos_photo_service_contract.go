package contracts

import (
	"context"

	"cariinkos/app/dto"
	"cariinkos/app/models"
)

type KosPhotoService interface {
	GetAllByKosID(kosID string) ([]models.KosPhoto, error)
	Create(ctx context.Context, userID, kosID string, request dto.CreateKosPhotoRequest) (models.KosPhoto, error)
	Update(ctx context.Context, userID, kosID, photoID string, request dto.UpdateKosPhotoRequest) (models.KosPhoto, error)
	Delete(ctx context.Context, userID, kosID, photoID string) error
}
