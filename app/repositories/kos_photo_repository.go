package repositories

import (
	"errors"

	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"

	"gorm.io/gorm"
)

type KosPhotoRepository struct{}

func NewKosPhotoRepository() *KosPhotoRepository { return &KosPhotoRepository{} }

func (r *KosPhotoRepository) GetAllByKosID(kosID string) ([]models.KosPhoto, error) {
	var photos []models.KosPhoto
	err := facades.Orm().Query().Where("kos_id = ?", kosID).Find(&photos)
	return photos, err
}

func (r *KosPhotoRepository) GetByID(kosID, photoID string) (models.KosPhoto, error) {
	var photo models.KosPhoto
	err := facades.Orm().Query().Where("kos_id = ? AND id = ?", kosID, photoID).First(&photo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.KosPhoto{}, helpers.ErrDataNotFound
	}
	return photo, err
}

func (r *KosPhotoRepository) Create(photo *models.KosPhoto) error {
	return facades.Orm().Query().Create(&photo)
}

func (r *KosPhotoRepository) Update(photo *models.KosPhoto) error {
	var tablename = photo.TableName() // get table name
	result, err := facades.Orm().Query().Table(tablename).Where("kos_id = ? AND id = ?", photo.KosID, photo.ID).Update(map[string]any{
		"image_url": photo.ImageURL,
		"public_id": photo.PublicID,
	})
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}
	return nil
}

func (r *KosPhotoRepository) Delete(kosID, photoID string) error {
	var photo models.KosPhoto
	result, err := facades.Orm().Query().Where("kos_id = ? AND id = ?", kosID, photoID).Delete(&photo)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}
	return nil
}

var _ repositoryContracts.KosPhotoRepositoryContract = (*KosPhotoRepository)(nil)
