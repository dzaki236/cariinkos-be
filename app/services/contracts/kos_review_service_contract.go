package contracts

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
)

type KosReviewService interface {
	GetByKos(kosID string) ([]models.KosReview, error)
	Create(userID string, kosID string, request dto.ReviewRequest) (models.KosReview, error)
	Update(userID string, kosID string, request dto.ReviewRequest) (models.KosReview, error)
	Delete(userID string, kosID string) error
}
