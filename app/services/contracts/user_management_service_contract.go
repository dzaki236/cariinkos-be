package contracts

import (
	"errors"

	"cariinkos/app/models"
)

var ErrNotManageableUser = errors.New("only users with role user can be managed")

type UserManagementService interface {
	GetUsers() ([]models.User, error)
	UpdateUserStatus(id string, active bool) (models.User, error)
}
