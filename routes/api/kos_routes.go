package api

import (
	"cariinkos/app/providers"

	"github.com/goravel/framework/contracts/route"
)

func KosRoutes(router route.Router) {
	controller := providers.NewKosProvider()
	router.Get("/kos", controller.Index)
	router.Get("/kos/{id}", controller.Show)
	router.Post("/kos", controller.Store)
	router.Put("/kos/{id}", controller.Update)
	router.Delete("/kos/{id}", controller.Destroy)
}
