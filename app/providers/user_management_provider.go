package providers

import (
	"cariinkos/app/http/controllers"
	"cariinkos/app/repositories"
	"cariinkos/app/services"
)

func NewUserManagementProvider() *controllers.UserManagementController {
	repository := repositories.NewUserManagementRepository()
	service := services.NewUserManagementService(repository)
	return controllers.NewUserManagementController(service)
}
