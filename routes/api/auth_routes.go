package api

import (
	"cariinkos/app/http/middleware"
	"cariinkos/app/providers"

	"github.com/goravel/framework/contracts/route"
)

func AuthRoutes(router route.Router) {
	controller := providers.NewUserAuthProvider()
	router.Post("users/login", controller.Login)
	router.Post("users/register", controller.Register)
	router.Middleware(middleware.NewHandleAuth()).Group(func(router route.Router) {
		router.Post("users/logout", controller.Logout)
	})
}
