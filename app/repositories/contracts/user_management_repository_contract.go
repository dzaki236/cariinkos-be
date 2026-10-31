package contracts

import "cariinkos/app/models"

type UserManagementRepository interface {
	FindByID(id string) (models.User, error)
	GetUsers() ([]models.User, error)
	Update(user *models.User) error
}
