package contracts

import "cariinkos/app/models"

type KosFacilityService interface {
	GetAllByKosID(kosID string) ([]models.Facility, error)
	Attach(kosID, facilityID string) error
	Detach(kosID, facilityID string) error
}
