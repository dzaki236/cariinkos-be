package services

import (
	"time"

	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"
)

type UserManagementService struct {
	repository repositoryContracts.UserManagementRepository
}

func NewUserManagementService(repository repositoryContracts.UserManagementRepository) *UserManagementService {
	return &UserManagementService{repository: repository}
}

func (s *UserManagementService) GetUsers() ([]models.User, error) {
	return s.repository.GetUsers()
}

func (s *UserManagementService) UpdateUserStatus(id string, active bool) (models.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return models.User{}, err
	}
	if user.Role != "user" {
		return models.User{}, serviceContracts.ErrNotManageableUser
	}

	if active {
		user.DisabledAt = nil
	} else {
		now := time.Now()
		user.DisabledAt = &now
	}

	if err := s.repository.Update(&user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

var _ serviceContracts.UserManagementService = (*UserManagementService)(nil)
