package contracts

import "cariinkos/app/models"

type KosFacilityRepositoryContract interface {
	GetAllByKosID(kosID string) ([]models.Facility, error)
	Attach(kosID, facilityID string) error
	Detach(kosID, facilityID string) error
}
