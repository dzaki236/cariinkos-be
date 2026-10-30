package providers

import (
	"cariinkos/app/http/controllers"
	"cariinkos/app/repositories"
	"cariinkos/app/services"
)

func NewKosFacilityProvider() *controllers.KosFacilityController {
	repository := repositories.NewKosFacilityRepository()
	service := services.NewKosFacilityService(repository)
	return controllers.NewKosFacilityController(service)
}
