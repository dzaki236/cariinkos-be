package bootstrap

import (
	"github.com/goravel/framework/contracts/database/seeder"

	"cariinkos/database/seeders"
)

func Seeders() []seeder.Seeder {
	return []seeder.Seeder{
		&seeders.UserSeeder{},
	}
}
