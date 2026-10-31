package services

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"
	"github.com/google/uuid"
)

type KosFavoriteService struct {
	repo repositoryContracts.KosFavoriteRepository
}

func NewKosFavoriteService(repo repositoryContracts.KosFavoriteRepository) *KosFavoriteService {
	return &KosFavoriteService{repo: repo}
}
func (s *KosFavoriteService) GetByUser(userID string) ([]models.KosFavorite, error) {
	return s.repo.GetByUser(userID)
}
func (s *KosFavoriteService) Create(userID, kosID string, request dto.UpdateFavoriteRequest) (models.KosFavorite, error) {
	favorite := models.KosFavorite{ID: uuid.NewString(), UserID: userID, KosID: kosID, Note: request.Note}
	if err := s.repo.Create(&favorite); err != nil {
		return models.KosFavorite{}, err
	}
	return favorite, nil
}
func (s *KosFavoriteService) Update(userID, kosID string, request dto.UpdateFavoriteRequest) (models.KosFavorite, error) {
	favorite, err := s.repo.Find(userID, kosID)
	if err != nil {
		return models.KosFavorite{}, err
	}
	favorite.Note = request.Note
	if err := s.repo.Update(&favorite); err != nil {
		return models.KosFavorite{}, err
	}
	return favorite, nil
}
func (s *KosFavoriteService) Delete(userID, kosID string) error { return s.repo.Delete(userID, kosID) }

var _ serviceContracts.KosFavoriteService = (*KosFavoriteService)(nil)
