package api

import (
	"cariinkos/app/providers"

	"github.com/goravel/framework/contracts/route"
)

func KosRoutes(router route.Router) {
	controller := providers.NewKosProvider()
	facilityController := providers.NewKosFacilityProvider()
	photoController := providers.NewKosPhotoProvider()
	contactController := providers.NewKosContactProvider()

	router.Get("/kos", controller.Index)
	router.Get("/kos/{id}", controller.Show)
	router.Post("/kos", controller.Store)
	router.Put("/kos/{id}", controller.Update)
	router.Delete("/kos/{id}", controller.Destroy)

	router.Get("/kos/{id}/facilities", facilityController.Index)
	router.Post("/kos/{id}/facilities", facilityController.Store)
	router.Delete("/kos/{id}/facilities/{facility_id}", facilityController.Destroy)

	router.Get("/kos/{id}/photos", photoController.Index)
	router.Post("/kos/{id}/photos", photoController.Store)
	router.Put("/kos/{id}/photos/{photo_id}", photoController.Update)
	router.Delete("/kos/{id}/photos/{photo_id}", photoController.Destroy)

	router.Get("/kos/{id}/contacts", contactController.Index)
	router.Post("/kos/{id}/contacts", contactController.Store)
	router.Put("/kos/{id}/contacts", contactController.Update)
	router.Delete("/kos/{id}/contacts", contactController.Destroy)
}
