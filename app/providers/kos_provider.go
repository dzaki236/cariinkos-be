package providers

import (
	"cariinkos/app/http/controllers"
	"cariinkos/app/repositories"
	"cariinkos/app/services"
)

func NewKosProvider() *controllers.KosController {
	repository := repositories.NewKosRepository()
	service := services.NewKosService(repository)
	return controllers.NewKosController(service)
}
