package providers

import (
	"cariinkos/app/http/controllers"
	"cariinkos/app/repositories"
	"cariinkos/app/services"
)

func NewKosContactProvider() *controllers.KosContactController {
	repository := repositories.NewKosContactRepository()
	service := services.NewKosContactService(repository)
	return controllers.NewKosContactController(service)
}
