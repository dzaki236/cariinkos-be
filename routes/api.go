package routes

import (
	"cariinkos/app/http/middleware"
	apiRoutes "cariinkos/routes/api"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Api() {
	facades.Route().Prefix("api/v1").Group(func(router route.Router) {
		router.Get("/", func(ctx http.Context) http.Response {
			return ctx.Response().Success().Json(http.Json{"halo": "api"})
		})
		router.Get("/docs", func(ctx http.Context) http.Response {
			return ctx.Response().View().Make("apidocs.tmpl", map[string]any{})
		})
		apiRoutes.AuthRoutes(router)
		router.Middleware(middleware.NewHandleAuth()).Group(func(router route.Router) {
			apiRoutes.KosRoutes(router)
			apiRoutes.FacilityRoutes(router)
		})
	})
}
