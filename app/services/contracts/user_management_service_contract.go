package contracts

import "cariinkos/app/models"

type UserManagementService interface {
	GetUsers() ([]models.User, error)
	UpdateUserStatus(id string, active bool) (models.User, error)
}
