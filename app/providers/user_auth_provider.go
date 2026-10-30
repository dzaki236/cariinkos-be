package providers

import (
	"cariinkos/app/http/controllers"
	"cariinkos/app/repositories"
	"cariinkos/app/services"
)

func NewUserAuthProvider() *controllers.AuthController {
	repository := repositories.NewUserAuthRepository()
	service := services.NewUserAuthService(repository)
	return controllers.NewAuthController(service)
}
