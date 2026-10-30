package services

import (
	"errors"

	"cariinkos/app/dto"
	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/goravel/framework/contracts/http"
)

type UserAuthService struct {
	repository repositoryContracts.UserAuthRepository
}

func NewUserAuthService(repository repositoryContracts.UserAuthRepository) *UserAuthService {
	return &UserAuthService{repository: repository}
}

func (s *UserAuthService) Login(ctx http.Context, request dto.LoginRequest) (models.User, string, error) {
	user, err := s.repository.FindByEmail(request.Email)
	if err != nil {
		if errors.Is(err, helpers.ErrDataNotFound) {
			return models.User{}, "", serviceContracts.ErrInvalidCredentials
		}
		return models.User{}, "", err
	}

	if !facades.Hash().Check(request.Password, user.Password) {
		return models.User{}, "", serviceContracts.ErrInvalidCredentials
	}

	token, err := facades.Auth(ctx).Login(user)
	if err != nil {
		return models.User{}, "", err
	}

	return user, token, nil
}

func (s *UserAuthService) Logout(ctx http.Context) error {
	return facades.Auth(ctx).Logout()
}

var _ serviceContracts.UserAuthService = (*UserAuthService)(nil)
