package contracts

import "cariinkos/app/models"

type FacilityRepositoryContract interface {
	GetAllData() ([]models.Facility, error)
	GetByID(id string) (models.Facility, error)
	Create(facility *models.Facility) error
	Update(facility *models.Facility) error
	Delete(id string) error
}
