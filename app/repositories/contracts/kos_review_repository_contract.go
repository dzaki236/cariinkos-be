package contracts

import "cariinkos/app/models"

type KosReviewRepository interface {
	GetByKos(kosID string) ([]models.KosReview, error)
	Find(userID, kosID string) (models.KosReview, error)
	Create(review *models.KosReview) error
	Update(review *models.KosReview) error
	Delete(userID, kosID string) error
}
