package services

import (
	"errors"

	"cariinkos/app/dto"
	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/google/uuid"
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
			return models.User{}, "", helpers.ErrInvalidCredentials
		}
		return models.User{}, "", err
	}

	if !facades.Hash().Check(request.Password, user.Password) {
		return models.User{}, "", helpers.ErrInvalidCredentials
	}

	token, err := facades.Auth(ctx).Login(user)
	if err != nil {
		return models.User{}, "", err
	}

	return user, token, nil
}

func (s *UserAuthService) Register(request dto.RegisterRequest) (models.User, error) {
	password, err := facades.Hash().Make(request.Password)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		ID:       uuid.NewString(),
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
		Role:     "user",
	}
	if err := s.repository.Create(&user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserAuthService) Logout(ctx http.Context) error {
	return facades.Auth(ctx).Logout()
}

func (s *UserAuthService) UpdateProfile(ctx http.Context, request dto.UpdateProfileRequest) (models.User, error) {
	userID, err := facades.Auth(ctx).ID()
	if err != nil {
		return models.User{}, err
	}

	user, err := s.repository.FindByID(userID)
	if err != nil {
		return models.User{}, err
	}

	user.Name = request.Name
	user.Email = request.Email
	if request.NewPassword != nil && *request.NewPassword != "" {
		if request.OldPassword == nil || !facades.Hash().Check(*request.OldPassword, user.Password) {
			return models.User{}, helpers.ErrInvalidProfilePassword
		}
		if request.NewPasswordConfirmation == nil || *request.NewPassword != *request.NewPasswordConfirmation {
			return models.User{}, helpers.ErrPasswordConfirmationMismatch
		}

		password, err := facades.Hash().Make(*request.NewPassword)
		if err != nil {
			return models.User{}, err
		}
		user.Password = password
	}
	if err := s.repository.UpdateProfile(&user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

var _ serviceContracts.UserAuthService = (*UserAuthService)(nil)
