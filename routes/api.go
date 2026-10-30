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
			return ctx.Response().Status(http.StatusCreated).Json(http.Json{"status": "ok"})
		})
		router.Get("/docs", func(ctx http.Context) http.Response {
			return ctx.Response().View().Make("apidocs.tmpl", map[string]any{})
		})
		apiRoutes.AuthRoutes(router)
		apiRoutes.KosRoutes(router)
		router.Middleware(middleware.NewHandleAuth()).Group(func(router route.Router) {
			apiRoutes.FacilityRoutes(router)
		})
	})
}
