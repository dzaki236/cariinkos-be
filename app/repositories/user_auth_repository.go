package repositories

import (
	"errors"

	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
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

	err := facades.Orm().Query().Where("email = ?", email).FirstOr(&user, func() error {
		return helpers.ErrDataNotFound
	})
	return user, err
}

func (r *UserAuthRepository) FindByID(id string) (models.User, error) {
	var user models.User
	err := facades.Orm().Query().Where("id = ?", id).FirstOr(&user, func() error {
		return helpers.ErrDataNotFound
	})
	return user, err
}

func (r *UserAuthRepository) UpdateProfile(user *models.User) error {
	var tablename string = user.TableName() // get table name
	_, err := facades.Orm().Query().Table(tablename).Where("id = ?", user.ID).Update(map[string]any{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
	return err
}

func (r *UserAuthRepository) Create(user *models.User) error {
	return facades.Orm().Query().Create(user)
}

var _ repositoryContracts.UserAuthRepository = (*UserAuthRepository)(nil)
