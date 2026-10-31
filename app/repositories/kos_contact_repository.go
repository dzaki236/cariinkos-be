package repositories

import (
	"errors"

	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"

	goravelErrors "github.com/goravel/framework/errors"
)

type KosContactRepository struct{}

func NewKosContactRepository() *KosContactRepository { return &KosContactRepository{} }

func (r *KosContactRepository) GetByKosID(kosID string) (models.KosContact, error) {
	var contact models.KosContact
	err := facades.Orm().Query().Where("kos_id = ?", kosID).First(&contact)
	if errors.Is(err, goravelErrors.OrmRecordNotFound) {
		return models.KosContact{}, helpers.ErrDataNotFound
	}
	return contact, err
}

func (r *KosContactRepository) Create(contact *models.KosContact) error {
	return facades.Orm().Query().Create(contact)
}

func (r *KosContactRepository) Update(contact *models.KosContact) error {
	_, err := facades.Orm().Query().Where("kos_id = ?", contact.KosID).Update(contact)
	return err
}

func (r *KosContactRepository) Delete(kosID string) error {
	var contact models.KosContact
	result, err := facades.Orm().Query().Where("kos_id = ?", kosID).Delete(&contact)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}
	return nil
}

var _ repositoryContracts.KosContactRepositoryContract = (*KosContactRepository)(nil)
