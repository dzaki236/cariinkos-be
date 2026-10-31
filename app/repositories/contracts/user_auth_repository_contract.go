package contracts

import "cariinkos/app/models"

type UserAuthRepository interface {
	FindByEmail(email string) (models.User, error)
	FindByID(id string) (models.User, error)
	Create(user *models.User) error
	UpdateProfile(user *models.User) error
}
