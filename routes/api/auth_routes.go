package api

import (
	"cariinkos/app/http/middleware"
	"cariinkos/app/providers"

	"github.com/goravel/framework/contracts/route"
)

func AuthRoutes(router route.Router) {
	controller := providers.NewUserAuthProvider()
	router.Post("/login", controller.Login)
	router.Middleware(middleware.NewHandleAuth()).Group(func(router route.Router) {
		router.Post("/logout", controller.Logout)
	})
}
