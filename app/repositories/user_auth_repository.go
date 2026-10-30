package repositories

import (
	"errors"

	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"

	"gorm.io/gorm"
)

type UserAuthRepository struct{}

func NewUserAuthRepository() *UserAuthRepository {
	return &UserAuthRepository{}
}

func (r *UserAuthRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	if email == "" {
		return user, errors.New("email is required")
	}

	err := facades.Orm().Query().Where("email = ?", email).First(&user)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, helpers.ErrDataNotFound
	}
	return user, err
}

func (r *UserAuthRepository) FindByID(id string) (models.User, error) {
	var user models.User
	err := facades.Orm().Query().Where("id = ?", id).First(&user)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, helpers.ErrDataNotFound
	}
	return user, err
}

var _ repositoryContracts.UserAuthRepository = (*UserAuthRepository)(nil)
