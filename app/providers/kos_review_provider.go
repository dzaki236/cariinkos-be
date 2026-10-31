package providers

import (
	"cariinkos/app/http/controllers"
	"cariinkos/app/repositories"
	"cariinkos/app/services"
)

func NewKosReviewProvider() *controllers.KosReviewController {
	return controllers.NewKosReviewController(services.NewKosReviewService(repositories.NewKosReviewRepository()))
}
