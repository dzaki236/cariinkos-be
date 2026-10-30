package api

import (
	"cariinkos/app/providers"

	"github.com/goravel/framework/contracts/route"
)

func FacilityRoutes(router route.Router) {
	controller := providers.NewFacilityProvider()
	router.Get("/facilities", controller.Index)
	router.Get("/facilities/{id}", controller.Show)
	router.Post("/facilities", controller.Store)
	router.Put("/facilities/{id}", controller.Update)
	router.Delete("/facilities/{id}", controller.Destroy)
}
