package repositories

import (
	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
)

type KosReviewRepository struct{}

func NewKosReviewRepository() *KosReviewRepository { return &KosReviewRepository{} }

func (r *KosReviewRepository) GetByKos(kosID string) ([]models.KosReview, error) {
	var reviews []models.KosReview
	err := facades.Orm().Query().Where("kos_id = ?", kosID).With("User").Order("created_at desc").Get(&reviews)
	return reviews, err
}

func (r *KosReviewRepository) Find(userID, kosID string) (models.KosReview, error) {
	var review models.KosReview
	err := facades.Orm().Query().Where("user_id = ? AND kos_id = ?", userID, kosID).With("User").FirstOr(&review, func() error {
		return helpers.ErrDataNotFound
	})
	return review, err
}

func (r *KosReviewRepository) Create(review *models.KosReview) error {
	return facades.Orm().Query().Create(&review)
}

func (r *KosReviewRepository) Update(review *models.KosReview) error {
	_, err := facades.Orm().Query().Where("id = ?", review.ID).Update(review)
	return err
}

func (r *KosReviewRepository) Delete(userID, kosID string) error {
	var review models.KosReview
	result, err := facades.Orm().Query().Where("user_id = ? AND kos_id = ?", userID, kosID).Delete(&review)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}
	return nil
}

var _ repositoryContracts.KosReviewRepository = (*KosReviewRepository)(nil)
