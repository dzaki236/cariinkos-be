package contracts

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
)

type KosReviewService interface {
	GetByKos(kosID string) ([]models.KosReview, error)
	Create(userID, kosID string, request dto.ReviewRequest) (models.KosReview, error)
	Update(userID, kosID string, request dto.ReviewRequest) (models.KosReview, error)
	Delete(userID, kosID string) error
}
