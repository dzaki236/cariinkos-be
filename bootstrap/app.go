package bootstrap

import (
	"cariinkos/config"
	"cariinkos/routes"
	contractsfoundation "github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/foundation"
)

func Boot() contractsfoundation.Application {
	return foundation.Setup().
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
