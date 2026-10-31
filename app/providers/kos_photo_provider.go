package providers

import (
	"cariinkos/app/helpers"
	"cariinkos/app/http/controllers"
	"cariinkos/app/repositories"
	"cariinkos/app/services"
)

func NewKosPhotoProvider() *controllers.KosPhotoController {
	uploader := helpers.NewCloudinaryPhotoUploader()
	repository := repositories.NewKosPhotoRepository()
	kosRepository := repositories.NewKosRepository()
	service := services.NewKosPhotoService(repository, kosRepository, uploader)
	return controllers.NewKosPhotoController(service)
}
