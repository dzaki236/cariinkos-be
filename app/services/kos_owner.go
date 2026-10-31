package services

import (
	"cariinkos/app/helpers"
	"cariinkos/app/repositories/contracts"
)

func CheckKosOwner(userID, kosID string, repository contracts.KosRepositoryContract) error {
	kos, err := repository.GetByID(kosID)
	if err != nil {
		return err
	}
	if kos.UserID != userID {
		return helpers.ErrDataNotFound
	}
	return nil
}
