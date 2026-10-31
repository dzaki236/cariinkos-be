package contracts

import "cariinkos/app/models"

type KosFacilityService interface {
	GetAllByKosID(kosID string) ([]models.Facility, error)
	Attach(userID string, kosID, facilityID string) error
	Detach(userID string, kosID, facilityID string) error
}
