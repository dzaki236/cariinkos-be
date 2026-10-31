package repositories

import (
	"errors"

	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"

	goravelErrors "github.com/goravel/framework/errors"
)

type UserManagementRepository struct{}

func NewUserManagementRepository() *UserManagementRepository {
	return &UserManagementRepository{}
}

func (r *UserManagementRepository) FindByID(id string) (models.User, error) {
	var user models.User
	err := facades.Orm().Query().Where("id = ?", id).First(&user)
	if errors.Is(err, goravelErrors.OrmRecordNotFound) {
		return models.User{}, helpers.ErrDataNotFound
	}
	return user, err
}

func (r *UserManagementRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := facades.Orm().Query().Where("role = ?", "user").Order("created_at desc").Get(&users)
	return users, err
}

func (r *UserManagementRepository) Update(user *models.User) error {
	_, err := facades.Orm().Query().Where("id = ?", user.ID).Update(user)
	return err
}

var _ repositoryContracts.UserManagementRepository = (*UserManagementRepository)(nil)
