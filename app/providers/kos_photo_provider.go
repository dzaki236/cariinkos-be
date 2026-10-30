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
	service := services.NewKosPhotoService(repository, uploader)
	return controllers.NewKosPhotoController(service)
}
