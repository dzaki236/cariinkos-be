package providers

import (
	"cariinkos/app/http/controllers"
	"cariinkos/app/repositories"
	"cariinkos/app/services"
)

func NewFacilityProvider() *controllers.FacilityController {
	repository := repositories.NewFacilityRepository()
	service := services.NewFacilityService(repository)
	return controllers.NewFacilityController(service)
}
