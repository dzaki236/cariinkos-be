package providers

import (
	"cariinkos/app/http/controllers"
	"cariinkos/app/repositories"
	"cariinkos/app/services"
)

func NewKosContactProvider() *controllers.KosContactController {
	repository := repositories.NewKosContactRepository()
	kosRepository := repositories.NewKosRepository()
	service := services.NewKosContactService(repository, kosRepository)
	return controllers.NewKosContactController(service)
}
