package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"cariinkos/app/facades"
)

type M20261029043717CreateUsersTable struct{}

// Signature The unique signature for the migration.
func (r *M20261029043717CreateUsersTable) Signature() string {
	return "20261029043717_create_users_table"
}

// Up Run the migrations.
func (r *M20261029043717CreateUsersTable) Up() error {
	if !facades.Schema().HasTable("users") {
		return facades.Schema().Create("users", func(table schema.Blueprint) {
			table.Uuid("id")
			table.String("name")
			table.String("email")
			table.String("password")
			table.String("role").Default("owner")
			table.Primary("id")
			table.Unique("email")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20261029043717CreateUsersTable) Down() error {
	return facades.Schema().DropIfExists("users")
}
