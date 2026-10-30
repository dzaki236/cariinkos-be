package bootstrap

import (
	"github.com/goravel/framework/contracts/database/schema"

	"cariinkos/database/migrations"
)

func Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20210101000001CreateJobsTable{},
		&migrations.M20261029043717CreateUsersTable{},
		&migrations.M20261029043803CreateKosTable{},
		&migrations.M20261029044803CreateFacilitiesTable{},
		&migrations.M20261029044809CreateKosFacilitiesTable{},
		&migrations.M20261029044814CreateKosContactsTable{},
		&migrations.M20261029044824CreateKosPhotosTable{},
	}
}
