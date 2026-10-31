package repositories

import (
	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	"errors"

	"gorm.io/gorm"
)

type KosFavoriteRepository struct{}

func NewKosFavoriteRepository() *KosFavoriteRepository { return &KosFavoriteRepository{} }

func (r *KosFavoriteRepository) GetByUser(userID string) ([]models.KosFavorite, error) {
	var favorites []models.KosFavorite
	err := facades.Orm().Query().Where("user_id = ?", userID).With("Kos").Order("created_at desc").Get(&favorites)
	return favorites, err
}

func (r *KosFavoriteRepository) Find(userID, kosID string) (models.KosFavorite, error) {
	var favorite models.KosFavorite
	err := facades.Orm().Query().Where("user_id = ? AND kos_id = ?", userID, kosID).With("Kos").First(&favorite)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.KosFavorite{}, helpers.ErrDataNotFound
	}
	return favorite, err
}

func (r *KosFavoriteRepository) Create(favorite *models.KosFavorite) error {
	return facades.Orm().Query().Create(favorite)
}

func (r *KosFavoriteRepository) Update(favorite *models.KosFavorite) error {
	_, err := facades.Orm().Query().Where("id = ?", favorite.ID).Update(favorite)
	return err
}

func (r *KosFavoriteRepository) Delete(userID, kosID string) error {
	var favorite models.KosFavorite
	result, err := facades.Orm().Query().Where("user_id = ? AND kos_id = ?", userID, kosID).Delete(&favorite)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}
	return nil
}

var _ repositoryContracts.KosFavoriteRepository = (*KosFavoriteRepository)(nil)
