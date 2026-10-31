package services

import (
	"errors"

	"cariinkos/app/dto"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/google/uuid"
)

type KosReviewService struct {
	repo repositoryContracts.KosReviewRepository
}

func NewKosReviewService(repo repositoryContracts.KosReviewRepository) *KosReviewService {
	return &KosReviewService{repo: repo}
}
func (s *KosReviewService) GetByKos(kosID string) ([]models.KosReview, error) {
	return s.repo.GetByKos(kosID)
}
func (s *KosReviewService) Create(userID, kosID string, request dto.ReviewRequest) (models.KosReview, error) {
	if _, err := s.repo.Find(userID, kosID); err == nil {
		return models.KosReview{}, helpers.ErrDataAlreadyExists
	} else if !errors.Is(err, helpers.ErrDataNotFound) {
		return models.KosReview{}, err
	}
	review := models.KosReview{ID: uuid.NewString(), UserID: userID, KosID: kosID, Rating: request.Rating, Comment: request.Comment}
	if err := s.repo.Create(&review); err != nil {
		return models.KosReview{}, err
	}
	return review, nil
}
func (s *KosReviewService) Update(userID, kosID string, request dto.ReviewRequest) (models.KosReview, error) {
	review, err := s.repo.Find(userID, kosID)
	if err != nil {
		return models.KosReview{}, err
	}
	review.Rating, review.Comment = request.Rating, request.Comment
	if err := s.repo.Update(&review); err != nil {
		return models.KosReview{}, err
	}
	return review, nil
}
func (s *KosReviewService) Delete(userID, kosID string) error { return s.repo.Delete(userID, kosID) }

var _ serviceContracts.KosReviewService = (*KosReviewService)(nil)
