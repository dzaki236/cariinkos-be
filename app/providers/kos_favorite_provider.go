package providers

import (
	"cariinkos/app/http/controllers"
	"cariinkos/app/repositories"
	"cariinkos/app/services"
)

func NewFavoriteProvider() *controllers.KosFavoriteController {
	return controllers.NewKosFavoriteController(services.NewKosFavoriteService(repositories.NewKosFavoriteRepository()))
}
