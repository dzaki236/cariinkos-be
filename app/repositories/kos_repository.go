package repositories

import (
	"errors"

	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"

	"gorm.io/gorm"
)

type KosRepository struct{}

func NewKosRepository() *KosRepository {
	return &KosRepository{}
}

func (r *KosRepository) GetAllData() ([]models.Kos, error) {
	var dataKos []models.Kos
	err := facades.Orm().Query().Find(&dataKos)
	return dataKos, err
}

func (r *KosRepository) GetByID(id string) (models.Kos, error) {
	var kos models.Kos
	err := facades.Orm().Query().Where("id = ?", id).First(&kos)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Kos{}, helpers.ErrDataNotFound
	}
	return kos, err
}

func (r *KosRepository) Create(kos *models.Kos) error {
	return facades.Orm().Query().Create(kos)
}

func (r *KosRepository) Update(kos *models.Kos) error {
	_, err := facades.Orm().Query().Where("id = ?", kos.ID).Update(&kos)
	return err
}

func (r *KosRepository) Delete(id string) error {
	var Kos models.Kos
	result, err := facades.Orm().Query().Where("id = ?", id).Delete(&Kos)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}
	return nil
}

var _ repositoryContracts.KosRepositoryContract = (*KosRepository)(nil)
