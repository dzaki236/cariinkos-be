package providers

import (
	"cariinkos/app/http/controllers"
	"cariinkos/app/repositories"
	"cariinkos/app/services"
)

func NewKosFacilityProvider() *controllers.KosFacilityController {
	repository := repositories.NewKosFacilityRepository()
	kosRepository := repositories.NewKosRepository()
	service := services.NewKosFacilityService(repository, kosRepository)
	return controllers.NewKosFacilityController(service)
}
