package repositories

import (
	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
)

type KosFacilityRepository struct{}

func NewKosFacilityRepository() *KosFacilityRepository {
	return &KosFacilityRepository{}
}

func (r *KosFacilityRepository) GetAllByKosID(kosID string) ([]models.Facility, error) {
	var facilities []models.Facility
	err := facades.Orm().Query().Table(models.Facility{}.TableName()).
		Join("JOIN kos_facilities ON kos_facilities.facility_id = facilities.id").
		Where("kos_facilities.kos_id = ?", kosID).
		Select("facilities.*").Find(&facilities)
	return facilities, err
}

func (r *KosFacilityRepository) Attach(kosID, facilityID string) error {
	var facility models.Facility
	if err := facades.Orm().Query().Where("id = ?", facilityID).FirstOr(&facility, func() error {
		return helpers.ErrDataNotFound
	}); err != nil {
		return err
	}
	return facades.Orm().Query().Create(&models.KosFacility{KosID: kosID, FacilityID: facilityID})
}

func (r *KosFacilityRepository) Detach(kosID, facilityID string) error {
	var pivot models.KosFacility
	result, err := facades.Orm().Query().Where("kos_id = ? AND facility_id = ?", kosID, facilityID).Delete(&pivot)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}
	return nil
}

var _ repositoryContracts.KosFacilityRepositoryContract = (*KosFacilityRepository)(nil)
