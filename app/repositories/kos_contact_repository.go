package repositories

import (
	"errors"

	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"

	"gorm.io/gorm"
)

type KosContactRepository struct{}

func NewKosContactRepository() *KosContactRepository { return &KosContactRepository{} }

func (r *KosContactRepository) GetAllByKosID(kosID string) ([]models.KosContact, error) {
	var contacts []models.KosContact
	err := facades.Orm().Query().Where("kos_id = ?", kosID).Find(&contacts)
	return contacts, err
}

func (r *KosContactRepository) GetByID(kosID, contactID string) (models.KosContact, error) {
	var contact models.KosContact
	err := facades.Orm().Query().Where("kos_id = ? AND id = ?", kosID, contactID).First(&contact)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.KosContact{}, helpers.ErrDataNotFound
	}
	return contact, err
}

func (r *KosContactRepository) Create(contact *models.KosContact) error {
	return facades.Orm().Query().Create(contact)
}

func (r *KosContactRepository) Update(contact *models.KosContact) error {
	_, err := facades.Orm().Query().Where("kos_id = ? AND id = ?", contact.KosID, contact.ID).Update(contact)
	return err
}

func (r *KosContactRepository) Delete(kosID, contactID string) error {
	var contact models.KosContact
	result, err := facades.Orm().Query().Where("kos_id = ? AND id = ?", kosID, contactID).Delete(&contact)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}
	return nil
}

var _ repositoryContracts.KosContactRepositoryContract = (*KosContactRepository)(nil)
