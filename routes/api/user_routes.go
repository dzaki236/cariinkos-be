package api

import (
	"cariinkos/app/http/middleware"
	"cariinkos/app/providers"

	"github.com/goravel/framework/contracts/route"
)

func UserRoutes(router route.Router) {
	controller := providers.NewUserAuthProvider()
	managementController := providers.NewUserManagementProvider()
	router.Middleware(middleware.NewHandleAuth()).Group(func(router route.Router) {
		router.Put("/users/profile", controller.Profile)
	})
	router.Middleware(middleware.NewHandleAuth(), middleware.NewRoleAuth("owner")).Group(func(router route.Router) {
		router.Get("/users", managementController.Index)
		router.Put("/users/{id}/status", managementController.UpdateStatus)
	})
}
