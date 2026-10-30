package repositories

import (
	"errors"

	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"

	"gorm.io/gorm"
)

type FacilityRepository struct{}

func NewFacilityRepository() *FacilityRepository {
	return &FacilityRepository{}
}

func (r *FacilityRepository) GetAllData() ([]models.Facility, error) {
	var facilities []models.Facility
	err := facades.Orm().Query().Find(&facilities)
	return facilities, err
}

func (r *FacilityRepository) GetByID(id string) (models.Facility, error) {
	var facility models.Facility
	err := facades.Orm().Query().Where("id = ?", id).First(&facility)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Facility{}, helpers.ErrDataNotFound
	}
	return facility, err
}

func (r *FacilityRepository) Create(facility *models.Facility) error {
	return facades.Orm().Query().Create(&facility)
}

func (r *FacilityRepository) Update(facility *models.Facility) error {
	_, err := facades.Orm().Query().Where("id = ?", facility.ID).Update(&facility)
	return err
}

func (r *FacilityRepository) Delete(id string) error {
	var Facility models.Facility
	result, err := facades.Orm().Query().Where("id = ?", id).Delete(&Facility)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}
	return nil
}

var _ repositoryContracts.FacilityRepositoryContract = (*FacilityRepository)(nil)
