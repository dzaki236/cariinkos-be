package contracts

import (
	"cariinkos/app/dto"
	"cariinkos/app/models"
)

type FacilityService interface {
	GetAllData() ([]models.Facility, error)
	GetByID(id string) (models.Facility, error)
	Create(request dto.CreateFacilityRequest) (models.Facility, error)
	Update(id string, request dto.UpdateFacilityRequest) (models.Facility, error)
	Delete(id string) error
}
