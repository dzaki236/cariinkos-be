package bootstrap

import (
	"cariinkos/app/facades"
	"cariinkos/app/http/middleware"
	"cariinkos/config"
	"cariinkos/routes"

	contractsfoundation "github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/contracts/foundation/configuration"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/foundation"
	"github.com/goravel/framework/http/limit"
)

func Boot() contractsfoundation.Application {
	return foundation.Setup().
		WithCallback(func() {
			facades.RateLimiter().For("global", func(ctx http.Context) http.Limit {
				return limit.PerDay(8804).Response(func(ctx http.Context) {
					ctx.Response().Status(http.StatusTooManyRequests).Json(http.Json{
						"message": "Rate limit exceeded. Please try again later.",
						"error":   "too many requests",
					}).Abort()
				})
			})
		}).
		WithMiddleware(func(m configuration.Middleware) {
			m.Append(middleware.NewRateLimiter())
		}).
		WithSeeders(Seeders).
		WithMigrations(Migrations).
		WithRouting(func() {
			routes.Web()
			routes.Grpc()
			routes.Api()
		}).
		WithProviders(Providers).
		WithConfig(config.Boot).
		Create()
}
